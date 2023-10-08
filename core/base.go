package core

import "github.com/gofiber/fiber/v2"

func BaseViews(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})
}
