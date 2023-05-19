package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"parser_labs/models"
)

func main() {
	app := fiber.New()

	groupV1 := app.Group("api/v1/")

	groupV1.Get("/parse/:name/:city/:key", models.GetAnalysis)

	log.Fatal(app.Listen(":9999"))
}
