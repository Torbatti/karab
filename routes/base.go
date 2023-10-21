package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Base(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"HeadTitle":   "کاراب",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})

}
