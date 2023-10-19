package core

import "github.com/gofiber/fiber/v2"

func Job(app *fiber.App) {
	app.Get("/jobs", func(c *fiber.Ctx) error {
		return c.Render("pages/jobs", fiber.Map{
			"HeadTitle":   "فرصت های شغلی",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})
}
