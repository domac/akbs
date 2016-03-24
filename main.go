package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/core"
	"github.com/phillihq/akbs/middleware"
	"net/http"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.Default()

	//设置中间件
	r.Use(middleware.MustParams)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "power by Gin")
	})

	go r.Run(":8080")

	//信号处理
	signalCH := core.InitSignal()
	core.HandleSignal(signalCH)
}
