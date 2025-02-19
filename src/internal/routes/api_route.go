package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nidrux/orca/internal/config"
)

const BASE_PATH string = "/api"

func ApiRoutes() {
	api := config.GetWebServer().Group(BASE_PATH)
	api.Get("/networks", func(context *fiber.Ctx) error {
		return context.Status(200).JSON(fiber.Map{"message": "Hello world"})
	})
	api.Get("/networks/:id", func(context *fiber.Ctx) error {
		return context.Status(200).JSON(fiber.Map{"message": "Hello world"})
	})
	api.Post("/networks", func(context *fiber.Ctx) error {
		return context.Status(200).JSON(fiber.Map{"message": "Hello world"})
	})
}
