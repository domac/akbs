package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {

		client := NewMyRedisClient(false)
		conn, err := client.GetConn()

		if err != nil {
			c.String(http.StatusBadRequest, "connect to redis fail")
			return
		}

		value := conn.Get("name").String()

		c.String(http.StatusOK, "power by Gin", value)
	})
}
