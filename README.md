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

## 感谢
煎鱼，学习来源https://book.eddycjy.com/golang/gin/install.html