package handler

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/core"
	"github.com/phillihq/akbs/logger"
	"net/http"
)

//Redis 链接测试处理
func RedisConnHandler(c *gin.Context) {
	logger.GetLogger().Infoln("Redis 连接测试 ")
	query := c.Query("query")
	cache, err := core.NewRedisCache()
	if err != nil {
		c.String(http.StatusOK, "power by Redigo cache err:"+err.Error())
	}
	val, err := redis.String(cache.Do("GET", query))
	if err != nil {
		c.String(http.StatusOK, "power by Redigo cache get err:"+err.Error())
	}

	c.String(http.StatusOK, "power by Redigo cache :"+val)
}
