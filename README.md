# gin-bin-example
学习来源https://book.eddycjy.com/golang/gin/install.html

## 安装与配置环境变量
> 安装

上官网找对应版本 https://studygolang.com/dl

> 环境变量

```
# Enable the go modules feature
export GO111MODULE=on
# Set the GOPROXY environment variable
export GOPROXY=https://goproxy.io
export GOROOT=/usr/local/go
export PATH=$PATH:GOROOT/bin
```

## Go Modules
go mod init [MODULE_PATH]：初始化 Go modules，它将会生成 go.mod 文件，需要注意的是 MODULE_PATH 填写的是模块引入路径，你可以根据自己的情况修改路径。

## Go 小技巧
> 重新整理go.mod indirect

```
$ go mod tidy
```
## 创建目录
conf：用于存储配置文件
middleware：应用中间件
models：应用数据库模型
pkg：第三方包
routers 路由逻辑处理
runtime：应用运行时数据

## Go Modules Replace
 go.mod 文件里,有Replace。首先你要看到我们使用的是完整的外部模块引用路径（github.com/fromChinaBoy/go-gin-example/xxx），而这个模块还没推送到远程，是没有办法下载下来的，因此需要用 replace 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。

## GORM
go get -u github.com/jinzhu/gorm

> 回调处理

gorm所支持的回调方法：
创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave

更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave

删除：BeforeDelete、AfterDelete

查询：AfterFind

## 常用库
- 验证: go get -u github.com/astaxie/beego/validation
- 常用工具: go get -u github.com/unknwon/com
- Orm: go get -u github.com/jinzhu/gorm
- Mysql: go get -u github.com/jinzhu/gorm/dialects/mysql
- Jwt: go get -u github.com/dgrijalva/jwt-go
- 配置管理: go get -u github.com/go-ini/ini
- 文档：go get -u github.com/swaggo/gin-swagger
- 文档：go get -u github.com/swaggo/gin-swagger/swaggerFiles
- 热更新：go get -u github.com/fvbock/endless
- 导入导出：go get -u github.com/360EntSecGroup-Skylar/excelize

## 热更新(需要unix系统)
安装endless: go get -u github.com/fvbock/endless

热更新：kill -1 pid

## Cron

## 感谢
Go by Example：https://gobyexample.com/
X分钟速成Y：https://learnxinyminutes.com/docs/zh-cn/go-cn/
煎鱼，学习来源: https://book.eddycjy.com/golang/gin/install.html
go和php的基础语法比较：https://engineering.carsguide.com.au/go-vs-php-syntax-comparison-c1465380b8ff