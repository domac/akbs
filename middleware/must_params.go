package middleware

import (
	"github.com/gin-gonic/gin"
)

//接口调用必须传入的参数
func MustParams(c *gin.Context) {
	app_key := c.Query("app_key")
	if len(app_key) == 0 {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}
