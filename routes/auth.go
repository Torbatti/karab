package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("pages/signup", fiber.Map{
			"HeadTitle":   "ثبت نام",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/login", fiber.Map{
			"HeadTitle":   "ورود",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})

}
