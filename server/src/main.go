package main

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/db"
	"comparisonLaboratories/src/herr"
	"comparisonLaboratories/src/transport"
	"github.com/go-pg/pg/v10"
	"os"
	"runtime"
)

func main() {
	clog.Logger.Info("Runtime", "Number of threads", runtime.GOMAXPROCS(runtime.NumCPU()-1))
	clog.Logger.Info("Runtime", "OS", runtime.GOOS, "Arch", runtime.GOARCH)

	core.InitEnv()
	core.InitServer(core.Server)
	core.InitConfig()
	transport.InitRouters(core.Server)

	err := db.ConnectToDB(&pg.Options{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("DATABASE"),
	})

	herr.HandlerError(err, "Fail connect to database")
	defer func(Database *pg.DB) {
		err := Database.Close()
		herr.HandlerError(err, "Unable to close database connection")

	}(db.Database)

	herr.HandlerError(core.Server.Run(), "Server did not start")
}
