package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/logger"
)

//Redis 链接测试处理
func RedisConnHandler(c *gin.Context) {
	logger.GetLogger().Infoln("Redis 连接测试 ")
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
}
