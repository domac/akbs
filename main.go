package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/core"
	"github.com/phillihq/akbs/logger"
	mdw "github.com/phillihq/akbs/middleware"
	router "github.com/phillihq/akbs/routes"
	"runtime"
)

var (
	debug      = flag.Bool("debug", false, "set debug mode")
	port       = *flag.String("port", "8080", "the server port")
	configFile = *flag.String("config", "./config/config.yaml", "the config file")
)

func main() {

	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	if *debug {
		//开启debug模式
		logger.GetLogger().Infoln("开启debug模式")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	conf, err := core.ParseConfigFile(configFile)

	if err != nil {
		panic("config file not found")
	}

	//设置配置信息
	core.SetConfig(conf)

	r := gin.New()

	//注册中间件
	mdw.RegisterMiddlewares(r)

	//注册路由
	router.RegisterRoutes(r)

	//运行web服务
	go r.Run(":" + port)

	//信号处理
	signalCH := core.InitSignal()
	core.HandleSignal(signalCH)
}
