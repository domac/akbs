package middleware

import (
	"github.com/gin-gonic/gin"
)

//注册中间件
func RegisterMiddlewares(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		CacheControl,
	)
}
