package main

import (
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/db"
	"comparisonLaboratories/src/transport"
	"github.com/go-pg/pg/v10"
	"log"
	"os"
)

func main() {
	core.InitEnv()
	core.InitServer(core.Server)
	core.InitConfig()
	transport.InitRouters(core.Server)

	err := db.ConnectToDB(&pg.Options{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("DATABASE"),
	})

	if err != nil {
		log.Fatalln(err)
		panic("database not connected")
	}
	defer func(Database *pg.DB) {
		err := Database.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db.Database)

	err = core.Server.Run()
	if err != nil {
		log.Fatalln("Server did not start")
	}
}
