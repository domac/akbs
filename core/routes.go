package core

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "power by Gin")
	})

	//测试redigo
	r.GET("/redigo", func(c *gin.Context) {
		rc, _ := OpenRedis("tcp", "192.168.139.139:6699", "")
		val, _ := redis.String(rc.Get().Do("GET", "name"))
		c.String(http.StatusOK, "power by Redigo:"+val)
	})
}
