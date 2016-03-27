package middleware

import (
	"github.com/gin-gonic/gin"
)

//设置访问缓存
func CacheControl(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "max-age=30, public, s-maxage=30")
}
