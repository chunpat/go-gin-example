package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/chunpat/go-gin-example/models"
	"github.com/chunpat/go-gin-example/pkg/gredis"
	"github.com/chunpat/go-gin-example/pkg/logging"
	"github.com/chunpat/go-gin-example/pkg/setting"
	"github.com/chunpat/go-gin-example/routers"
	"github.com/fvbock/endless"
)

func main() {
	//初始化配置
	setting.Setup()
	logging.SetUp()
	models.Setup()
	gredis.Setup()

	logging.Info("start server")

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
