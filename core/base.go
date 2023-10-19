package core

import (
	"github.com/gofiber/fiber/v2"
)

func Base(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"HeadTitle":   "Karab",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})

}
