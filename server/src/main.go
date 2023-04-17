package main

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/herr"
	"comparisonLaboratories/src/model/labs"
	"comparisonLaboratories/src/redis"
	"comparisonLaboratories/src/transport"
	"runtime"
)

func main() {
	clog.Logger.Info("Runtime", "Number of threads", runtime.GOMAXPROCS(runtime.NumCPU()-1))
	clog.Logger.Info("Runtime", "OS", runtime.GOOS, "Arch", runtime.GOARCH)

	core.InitEnv()
	labs.InitCities()
	core.InitServer(core.Server)
	core.InitConfig()
	redis.InitRedis()
	transport.InitRouters(core.Server)

	herr.HandlerError(core.Server.Run(), "Server did not start")
}
