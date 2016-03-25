package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/core"
	"github.com/phillihq/akbs/logger"
	"github.com/phillihq/akbs/middleware"
	"runtime"
)

var debug = flag.Bool("debug", false, "set debug mode")

func main() {

	if *debug {
		//开启debug模式
		logger.GetLogger().Infoln("开启debug模式")
		gin.SetMode(gin.DebugMode)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	r := gin.Default()

	//设置中间件
	r.Use(middleware.MustParams)

	core.RegisterRoutes(r)

	go r.Run(":8080")

	//信号处理
	signalCH := core.InitSignal()
	core.HandleSignal(signalCH)
}
