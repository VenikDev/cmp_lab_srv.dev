package main

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/core"
	"cmp_lab/src/herr"
	"cmp_lab/src/model/labs"
	"cmp_lab/src/redis"
	"cmp_lab/src/transport"
	"runtime"
)

func main() {
	clog.Logger.Info("[main/runtime]", "Number of threads", runtime.GOMAXPROCS(runtime.NumCPU()-1))
	clog.Logger.Info("[main/runtime]", "OS", runtime.GOOS, "Arch", runtime.GOARCH)

	core.InitEnv()
	city.InitCities()
	core.InitServer(core.Server)
	core.InitConfig()
	redis.InitRedis()
	transport.InitRouters(core.Server)

	err := core.Server.Run(":8080")
	if err != nil {
		herr.HandlerError(err, "Server did not start")
	}
}
