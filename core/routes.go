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

	r.GET("/cache", func(c *gin.Context) {
		query := c.Query("query")
		cache, err := NewRedisCache()
		if err != nil {
			c.String(http.StatusOK, "power by Redigo cache err:"+err.Error())
		}
		val, err := redis.String(cache.Do("GET", query))
		if err != nil {
			c.String(http.StatusOK, "power by Redigo cache get err:"+err.Error())
		}

		c.String(http.StatusOK, "power by Redigo cache :"+val)
	})
}
