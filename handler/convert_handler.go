package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConvertHandler(c *gin.Context) {
	name := c.PostForm("name")
	println(name)
	c.String(http.StatusOK, name)
}
