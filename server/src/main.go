package main

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/core"
	"cmp_lab/src/herr"
	"cmp_lab/src/model/labs"
	"cmp_lab/src/redis"
	"cmp_lab/src/transport"
	"os"
	"runtime"
)

func main() {
	clog.Logger.Info("[main/runtime]", "Number of threads", runtime.GOMAXPROCS(runtime.NumCPU()-1))
	clog.Logger.Info("[main/runtime]", "OS", runtime.GOOS, "Arch", runtime.GOARCH)

	core.InitEnv()
	labs.InitCities()
	core.InitServer(core.Server)
	core.InitConfig()
	redis.InitRedis()
	transport.InitRouters(core.Server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	clog.Logger.Info("[core/port]", "port", port)
	herr.HandlerError(core.Server.Run(":"+port), "Server did not start")
}
