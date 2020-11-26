
<img align="right" width="100" src="https://resource.fastgoo.net/go.png"/>

<h1 align="left"><a href="javascript:">gin脚手架模板</a></h1>

[![Latest Version](https://img.shields.io/badge/release-master-green.svg?maxAge=2592000)](https://github.com/fastgoo/gin-cli)
[![Golang Version](https://img.shields.io/badge/go-%3E=1.13-brightgreen.svg?maxAge=2592000)](https://studygolang.com/dl)
[![Gin Version](https://img.shields.io/badge/gin-%3E=1.6-brightgreen.svg?maxAge=2592000)](https://github.com/gin-gonic/gin)
[![Gorm Version](https://img.shields.io/badge/gorm-%3E=v2.0-brightgreen.svg?maxAge=2592000)](https://gorm.io/zh_CN/docs/index.html)
[![Logrus Version](https://img.shields.io/badge/logrus-%3E=v1.7-brightgreen.svg?maxAge=2592000)](https://github.com/sirupsen/logrus)

gin web框架脚手架服务，企业级快速构建web项目，不需要关心代码目录的结构以及常用方法的封装。

> gin 中文文档：https://learnku.com/docs/gin-gonic/2019
> gorm 中文文档：https://gorm.io/zh_CN/docs/index.html


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

## 支持功能
* [x] 服务平滑重启
* [x] gorm模型主从、单机
* [x] 自定义env配置文件
* [x] redis集群、单机的相关操作
* [x] 日志写入，日期分割
* [x] 自定义错误状态码、错误中英文文案、自定义错误类型
* [x] jwt授权认证
* [x] web请求日志记录
* [ ] casbin权限管理


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


## 目录结构
```shell
├── Makefile # 构建脚本
├── README.md # readme
├── go.mod # go mod
├── logs # 系统自动生成的日志目录
│   ├── test.log # 日志文件
├── main.go # 入口文件(启动gin服务，载入插件以及依赖类库)
├── middleware # 中间件
│   ├── casbin # 权限中间件
│   ├── gin # gin框架系统中间件
│   │   └── logger.go # 访问日志中间件
│   └── jwt # jwt授权验证中间件
│       └── jwt.go
├── modules # 业务模块(可建立多业务多模块)
│   └── api_v1
│       ├── api # api控制器方法
│       │   └── auth.go # 注册范例
│       ├── models # 模型
│       │   ├── model.go # 基础模型(处理连接池连接实例)
│       │   ├── model_base.go # gen生成的模型方法依赖
│       │   ├── wk_company.go # 范例1
│       │   └── wk_user_info.go
│       ├── module.go # 业务模块的核心(路由、模型初始化、验证器初始化等等)
│       ├── services # 业务逻辑
│       └── validator # 验证器
│           ├── auth.go # 范例验证器(注册的一系列操作)
│           └── init.go # 初始化验证器(每新增一个验证方法都需要，在这里手动注册引入)
├── pkg # 公共类
│   ├── e # 异常、消息、状态码
│   │   ├── code.go # 状态码
│   │   ├── error.go # 自定义错误结构
│   │   └── msg.go # 状态码指定的消息内容(需要手动定义，支持中英文)
│   ├── response # 响应类
│   │   └── response.go # 主要是处理http响应结果的(封装了错误、成功状态的处理机制，以及消息文字的处理)
│   └── util # 工具类(编写一些常用的方法给外部调用)
│       └── util.go
├── plugin.go # 插件类(目前没啥卵用)
├── plugins
│   ├── env # 处理.env配置文件的类
│   │   └── env.go
│   ├── logger # 日志类（在第三方类的基础之上做了封装，让写可控，目前只支持写入到文件）
│   │   ├── logger.go
│   │   ├── logrus.go
│   │   └── zap.go
│   └── redis # redis类 (支持单机和集群，内部预设了大部分常用方法，如果想自己添加直接在里面新增方法即可)
│       └── redis.go
├── router # 路由类（这里是吧gin的router抽离出来，让每一个module去引入使用，这样在模块删除的时候不会对外部造成影响）
│   └── router.go
└── end
```



## License

MIT
