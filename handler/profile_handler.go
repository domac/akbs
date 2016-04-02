package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/akbs/logger"
	"net/http"
	"runtime"
	"runtime/pprof"
)

func ProfileHandler(c *gin.Context) {

	name := c.Param("name")
	logger.GetLogger().Infoln("cmd:" + name)

	var profile *pprof.Profile

	switch name {
	case "goroutine":
		profile = pprof.Lookup("goroutine")
	case "heap":
		profile = pprof.Lookup("heap")
	case "block":
		profile = pprof.Lookup("block")
	case "threadcreate":
		profile = pprof.Lookup("threadcreate")
	case "mem":
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		c.JSON(http.StatusOK, gin.H{
			"mem alloc":       mem.Alloc,
			"mem total alloc": mem.TotalAlloc,
			"mem heap alloc":  mem.HeapAlloc,
			"mem heap sys":    mem.HeapSys,
		})
	default:
		logger.GetLogger().Warnln("no cmd")
	}
	if profile != nil {
		profile.WriteTo(c.Writer, 2)
	}
	c.String(http.StatusOK, "profile info")
}
