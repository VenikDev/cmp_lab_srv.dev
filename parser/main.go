package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"parser_labs/utils"
)

func main() {
	app := fiber.New()

	groupV1 := app.Group("api/v1/")

	groupV1.Get("/parse/:name/:city/:key", func(ctx *fiber.Ctx) error {
		callback := func() {
			ctx.SendStatus(fiber.StatusNotFound)
		}

		name := ctx.Params("name")
		utils.AssertCallback(name == "", callback)

		city := ctx.Params("city")
		utils.AssertCallback(city == "", callback)

		key := ctx.Params("key")
		utils.AssertCallback(key == "", callback)

		return ctx.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":9999"))
}
