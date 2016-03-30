package core

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "power by AKBS !")
	})

	r.GET("/cache", func(c *gin.Context) {
		query := c.Query("query")
		cache, err := NewRedisCache()
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
		}
		val, _ := redis.String(cache.Do("GET", query))
		c.String(http.StatusOK, "AKBS RESULTS :"+val)
	})
}
