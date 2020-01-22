package main

import (
	"fmt"
	"net/http"

	"github.com/FromChinaBoy/go-gin-example/pkg/setting"
	"github.com/FromChinaBoy/go-gin-example/routers"
)

func main() {
	var x int
	// 和if一样，for也不用括号
	for x = 0; x < 3; x++ { // ++ 自增
		fmt.Println("遍历", x)
	}
	fmt.Println("遍历", x)
	// 闭包函数
	xBig := func() bool {
		return x > 100 // x是上面声明的变量引用
	}
	fmt.Println("xBig:", xBig()) // true （上面把y赋给x了）

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
