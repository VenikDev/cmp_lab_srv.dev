package models

import (
	"github.com/gofiber/fiber/v2"
	"parser_labs/controllers"
	"parser_labs/models/store"
	"parser_labs/utils"
)

func GetAnalysis(ctx *fiber.Ctx) error {
	callback := func() {
		ctx.SendStatus(fiber.StatusNotFound)
	}

	name := ctx.Params("name")
	utils.AssertCallback(name == "", callback)

	city := ctx.Params("city")
	utils.AssertCallback(city == "", callback)

	key := ctx.Params("key")
	utils.AssertCallback(key == "", callback)

	params := make(store.StrStore, 3)
	params["name"] = name
	params["city"] = city
	params["key"] = key

	err := controllers.ProcessRequest(params)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
