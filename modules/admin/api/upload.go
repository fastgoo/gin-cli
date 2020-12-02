package api

import (
	"gin-cli/modules/admin/models"
	"gin-cli/pkg/e"
	"gin-cli/pkg/response"
	"gin-cli/pkg/util"
	"gin-cli/plugins/qiniu"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// 上传图片
func Upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, 200, e.ERR_UPLOAD_FILE_NOTFOUND)
		return
	}
	qiniuModel := models.WkQiniuFileRecord{
		UserID: 0,
		Name:   f.Filename,
		Size:   f.Size,
		Key:    "license/" + strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(util.GetRandom(1000, 9999)) + "." + util.FileExt(f.Filename),
		Hash:   "",
	}
	buf := make([]byte, f.Size)
	file, err := f.Open()
	file.Read(buf)
	err, hashKey, url := qiniu.UploadByBytes(buf, qiniuModel.Key)
	//c.SaveUploadedFile(f, "./"+f.Filename)
	//err, hashKey, url := qiniu.UploadByForm("./"+f.Filename, "test.jpg")
	if err != nil {
		response.Fail(c, 200, e.ERR_UPLOAD_FILE_QINIU_FAIL)
		return
	}
	qiniuModel.Hash = hashKey
	qiniuModel.Insert(qiniuModel)
	response.Success(c, e.SUCCESS, map[string]string{"url": url, "hash": hashKey})
}
