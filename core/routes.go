package core

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/logger"
	"net/http"
	"strconv"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {

		client := NewMyRedisClient(false)
		conn, err := client.GetConn()

		//defer conn.Close()

		fc := conn.PoolStats().FreeConns

		logger.GetLogger().Infof("fc : %s", strconv.Itoa(int(fc)))
		logger.GetLogger().Infoln(strconv.Itoa(int(conn.PoolStats().Waits)))

		if err != nil {
			c.String(http.StatusBadRequest, "connect to redis fail")
			return
		}

		value := conn.Get("name").String()

		c.String(http.StatusOK, "power by Gin", value)
	})
}
