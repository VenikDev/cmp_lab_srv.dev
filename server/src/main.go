package main

import (
	"comparisonLaboratories/src/db"
	"comparisonLaboratories/src/server"
	"comparisonLaboratories/src/transport"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.SetupServer(server.Server)
	transport.SetupRouters(server.Server)

	err = db.ConnectToDB(&pg.Options{
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

	users, err := db.GetCites(db.Database)
	if err != nil {
		log.Println(err)
	}
	log.Println(users)

	err = server.Server.Run()
	if err != nil {
		log.Fatalln("Server did not start")
	}
}
