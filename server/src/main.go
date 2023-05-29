package main

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/core"
	"cmp_lab/src/herr"
	"cmp_lab/src/redis"
	"cmp_lab/src/transport"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"runtime"
	"time"
)

func createLogFile() *os.File {
	apochNow := time.Now().Unix()
	logFile, err := os.Create(fmt.Sprintf("logs/%d.log", apochNow))
	if err != nil {
		err := os.Mkdir("logs", 0777)
		if err != nil {
			clog.Error("[main/create_dir]", "Failed create dir err", err)
			return nil
		}
	}

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	clog.Info("[main/log_file]", "create file is", "success")

	return logFile
}

func main() {
	clog.Info("[main/runtime]", "Number of threads", runtime.GOMAXPROCS(runtime.NumCPU()-1))
	clog.Info("[main/runtime]", "OS", runtime.GOOS, "Arch", runtime.GOARCH)

	logFile := createLogFile()
	defer logFile.Close()

	core.InitEnv()
	core.InitServer(core.Server)
	core.InitConfig()
	redis.InitRedis()
	transport.InitRouters(core.Server)

	err := core.Server.Run(":8080")
	if err != nil {
		herr.HandlerError(err, "Server did not start")
	}
}
