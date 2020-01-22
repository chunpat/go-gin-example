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
 go.mod 文件里,有Replace。首先你要看到我们使用的是完整的外部模块引用路径（github.com/EDDYCJY/go-gin-example/xxx），而这个模块还没推送到远程，是没有办法下载下来的，因此需要用 replace 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。

## GORM
go get -u github.com/jinzhu/gorm

> 回调处理

gorm所支持的回调方法：
创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind

## 常用组件


## 感谢
煎鱼，学习来源https://book.eddycjy.com/golang/gin/install.html