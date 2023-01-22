package main

import (
	"comparisonLaboratories/src/db"
	"comparisonLaboratories/src/server"
	"github.com/gin-gonic/gin"
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

	app := gin.Default()
	server.SetupServer(app)
	server.SetupRouters(app)

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

	err = app.Run()
	if err != nil {
		log.Fatalln("Server did not start")
	}
}
