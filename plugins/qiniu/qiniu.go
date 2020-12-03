package qiniu

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// 参考文档 https://github.com/qiniu/api.v7/wiki#public-get
type config struct {
	AccessKey string `env:"QINIU_ACCESS_KEY" envDefault:""`
	SecretKey string `env:"QINIU_SECRET_KEY" envDefault:""`
	Bucket    string `env:"QINIU_BUCKET" envDefault:""`
	Host      string `env:"QINIU_HOST" envDefault:""`
	Zone      string `env:"QINIU_ZONE" envDefault:"huadong"`
}

var qCfg = &config{}

type ProgressRecord struct {
	Progresses []storage.BlkputRet `json:"progresses"`
}

func init() {
	if err := env.Parse(qCfg); err != nil {
		log.Fatal(err)
	}
}

func md5Hex(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 获取客户端上传凭证
func GetUploadToken() string {
	putPolicy := storage.PutPolicy{
		Scope:   qCfg.Bucket,
		Expires: 7200, //有效时间可以自定义. 默认3600一个小时
	}
	mac := qbox.NewMac(qCfg.AccessKey, qCfg.SecretKey)
	return putPolicy.UploadToken(mac)
}

// 初始化基础存储配置信息
func newStorageConfig() (upToken string, cfg storage.Config) {
	upToken = GetUploadToken()
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	switch qCfg.Zone {
	case "huadong":
		cfg.Zone = &storage.ZoneHuadong
		break
	case "huabei":
		cfg.Zone = &storage.ZoneHuabei
		break
	case "huanan":
		cfg.Zone = &storage.ZoneHuanan
		break
	}
	return
}

func GetUrl(key string) string {
	return qCfg.Host + "/" + key
}

func UploadByForm(localFile, key string) (error, string, string) {
	token, cfg := newStorageConfig()
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutFile(context.Background(), &ret, token, key, localFile, nil)
	if err != nil {
		return err, "", ""
	}
	return nil, ret.Hash, qCfg.Host + "/" + key
}

func UploadByBytes(fileBytes []byte, key string) (error, string, string) {
	token, cfg := newStorageConfig()
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(context.Background(), &ret, token, key, bytes.NewReader(fileBytes), int64(len(fileBytes)), nil)
	if err != nil {
		return err, "", ""
	}
	return nil, ret.Hash, qCfg.Host + "/" + key
}

func UploadBySlice() {
	key := ""
	localFile := "/Users/jemy/Documents/github.png"
	token := GetUploadToken()
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{

	}
	err := resumeUploader.PutFile(context.Background(), &ret, token, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

// 断点续传
func UploadByContinue() {
	localFile := "your local file path"
	key := "your file save key"

	token := GetUploadToken()

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 必须仔细选择一个能标志上传唯一性的 recordKey 用来记录上传进度
	// 我们这里采用 md5(bucket+key+local_path+local_file_last_modified)+".progress" 作为记录上传进度的文件名
	fileInfo, statErr := os.Stat(localFile)
	if statErr != nil {
		fmt.Println(statErr)
		return
	}

	fileSize := fileInfo.Size()
	fileLmd := fileInfo.ModTime().UnixNano()
	recordKey := md5Hex(fmt.Sprintf("%s:%s:%s:%s", qCfg.Bucket, key, localFile, fileLmd)) + ".progress"

	// 指定的进度文件保存目录，实际情况下，请确保该目录存在，而且只用于记录进度文件
	recordDir := "/tmp/qiniu-progress"
	mErr := os.MkdirAll(recordDir, 0755)
	if mErr != nil {
		fmt.Println("mkdir for record dir error,", mErr)
		return
	}

	recordPath := filepath.Join(recordDir, recordKey)

	progressRecord := ProgressRecord{}
	// 尝试从旧的进度文件中读取进度
	recordFp, openErr := os.Open(recordPath)
	if openErr == nil {
		progressBytes, readErr := ioutil.ReadAll(recordFp)
		if readErr == nil {
			mErr := json.Unmarshal(progressBytes, &progressRecord)
			if mErr == nil {
				// 检查context 是否过期，避免701错误
				for _, item := range progressRecord.Progresses {
					if storage.IsContextExpired(item) {
						fmt.Println(item.ExpiredAt)
						progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
						break
					}
				}
			}
		}
		recordFp.Close()
	}

	if len(progressRecord.Progresses) == 0 {
		progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
	}

	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	progressLock := sync.RWMutex{}

	putExtra := storage.RputExtra{
		Progresses: progressRecord.Progresses,
		Notify: func(blkIdx int, blkSize int, ret *storage.BlkputRet) {
			progressLock.Lock()
			progressLock.Unlock()

			//将进度序列化，然后写入文件
			progressRecord.Progresses[blkIdx] = *ret
			progressBytes, _ := json.Marshal(progressRecord)
			fmt.Println("write progress file", blkIdx, recordPath)
			wErr := ioutil.WriteFile(recordPath, progressBytes, 0644)
			if wErr != nil {
				fmt.Println("write progress file error,", wErr)
			}
		},
	}
	err := resumeUploader.PutFile(context.Background(), &ret, token, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	//上传成功之后，一定记得删除这个进度文件
	os.Remove(recordPath)
	fmt.Println(ret.Key, ret.Hash)
}
