package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/middleware"
	"net/http"
)

func main() {
	r := gin.Default()

	//设置中间件
	r.Use(middleware.MustParams)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "power by Gin")
	})

	r.Run(":8080")
}
