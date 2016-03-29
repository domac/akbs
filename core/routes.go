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
		rc, err := OpenRedis("tcp", "192.168.139.139:7000", "")

		if err != nil {
			c.String(http.StatusOK, "power by Redigo err:"+err.Error())
		}

		val, err := redis.String(rc.Get().Do("GET", "name"))

		if err != nil {
			c.String(http.StatusOK, "power by Redigo get err:"+err.Error())
		}

		c.String(http.StatusOK, "power by Redigo:"+val)
	})

	r.GET("/cluster", func(c *gin.Context) {
		rc, err := OpenRedisCluster([]string{"192.168.139.139:7000"})

		if err != nil {
			c.String(http.StatusOK, "power by Redigo cluster err:"+err.Error())
		}

		val, err := redis.String(rc.Do("GET", "love"))

		if err != nil {
			c.String(http.StatusOK, "power by Redigo cluster get err:"+err.Error())
		}

		c.String(http.StatusOK, "power by Redigo cluster :"+val)
	})
}
