package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/core"
	"github.com/phillihq/akbs/logger"
	"net/http"
)

func MySQLConnHandler(c *gin.Context) {

	name := c.Param("name")

	session := core.OpenConnection()
	defer session.Close()

	type App struct {
		Id    string
		Token string
	}

	var apps []App
	session.Select(&apps, "select * from application where name = '"+name+"'")
	var token string
	var app App
	for _, app = range apps {
		token = app.Token
		logger.GetLogger().Infof("app token : %s \n", token)
	}
	c.String(http.StatusOK, "power by Mysql cache :"+token)
}
