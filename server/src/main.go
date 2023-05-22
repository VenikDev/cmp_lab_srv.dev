package main

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/core"
	"cmp_lab/src/herr"
	"cmp_lab/src/model/labs"
	"cmp_lab/src/redis"
	"cmp_lab/src/transport"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"runtime"
)

func main() {
	clog.Logger.Info("[main/runtime]", "Number of threads", runtime.GOMAXPROCS(runtime.NumCPU()-1))
	clog.Logger.Info("[main/runtime]", "OS", runtime.GOOS, "Arch", runtime.GOARCH)

	// Logging to a file.
	logFile, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile)

	core.InitEnv()
	labs.InitCities()
	core.InitServer(core.Server)
	core.InitConfig()
	redis.InitRedis()
	transport.InitRouters(core.Server)

	err := core.Server.Run(":8080")
	if err != nil {
		herr.HandlerError(err, "Server did not start")
	}
}
