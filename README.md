
<img align="right" width="100" src="https://resource.fastgoo.net/go.png"/>

<h1 align="left"><a href="javascript:">gin脚手架模板</a></h1>

[![Latest Version](https://img.shields.io/badge/release-master-green.svg?maxAge=2592000)](https://github.com/fastgoo/gin-cli)
[![Golang Version](https://img.shields.io/badge/go-%3E=1.13-brightgreen.svg?maxAge=2592000)](https://studygolang.com/dl)
[![Gin Version](https://img.shields.io/badge/gin-%3E=1.6-brightgreen.svg?maxAge=2592000)](https://github.com/gin-gonic/gin)
[![Gorm Version](https://img.shields.io/badge/gorm-%3E=v2.0-brightgreen.svg?maxAge=2592000)](https://gorm.io/zh_CN/docs/index.html)
[![Logrus Version](https://img.shields.io/badge/logrus-%3E=v1.7-brightgreen.svg?maxAge=2592000)](https://github.com/sirupsen/logrus)

gin web框架脚手架服务，企业级快速构建web项目，不需要关心代码目录的结构以及常用方法的封装。

> gin 中文文档：https://learnku.com/docs/gin-gonic/2019


## 安装

```shell
# 克隆整个项目
$ git clone https://github.com/fastgoo/gin-cli.git
$ cd gin-cli
# 设置系统配置
$ cp .env.example .env
# 调试运行服务
$ go run mai.go
# 编译构建
$ make
```

## 配置文件
刚拷贝下来的项目是没有`.env`的配置文件的，需要把`.env.example`改为`.env`
```ini
# 目前仅支持一主多从的数据库集群
# 主服务器的配置
MYSQL_MASTER=root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
# 从服务器的配置，如果是多个服务器，需要给字打引号，同时每个服务器之间需要使用,分割开来
MYSQL_SLAVES="root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local,root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
# 连接池中空闲连接可用时间
MYSQL_IDLETIME=3600
# 连接可复用的最大时间
MYSQL_LIFTTIME=86400
# 连接池中空闲连接的最大数量
MYSQL_IDLECONNS=20
# 数据库连接的最大数量
MYSQL_OPENCONNS=100

# 如果是redis集群配置，每个服务器之间需要使用,分割开来
REDIS_ADDRS="127.0.0.1:6379"
# redis密码
REDIS_PW=""
# 数据库，如果是集群那么这个设置选项无效(redis集群没有多数据库，只有db0)
REDIS_DB=0

# gin框架的运行模式 debug可以在控制台看到访问记录，prod则是在线上运行的时候设置 
GIN_MODE=debug
# api的语言，en英文，ch中文. 都需要提前预设好中英文的文本信息
LANGUAGE=ch
# gin 请求日志，记录在日志文件. 请求时间、请求ip、请求路由、响应时间等等信息
GIN_REQUEST_LOG=true
# web服务绑定ip，如果是外网访问那么这里为空就可以了，如果是本地访问那么就需要填写127.0.0.1
HTTP_HOST=""
# web服务监听端口
HTTP_PORT=9010

# jwt认证的秘钥
JWT_SECRET="test"

# 是否开启日志的写入功能
LOGGER_ENABLE=true
# 日志的写入文件，一般系统会自动根据时间来做分割
LOGGER_FILE="./logs/test.log"

```



## License

MIT
