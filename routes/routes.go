package core

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/handler"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "power by AKBS !")
	})
	r.GET("/cache", handler.RedisConnHandler) //测试redis

	r.GET("/db", handler.MySQLConnHandler) //测试mysql

	r.GET("/profile/:name", handler.ProfileHandler) //测试profile
}
