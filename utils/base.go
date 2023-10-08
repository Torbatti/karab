package utils

import "github.com/gofiber/fiber/v2"

func Stack(app *fiber.App) {
	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})
}
