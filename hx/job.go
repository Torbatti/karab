package hx

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// /hx/jobs/{{.ID}}

func Job(app *fiber.App) {
	app.Post("/hx/jobs/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			// TODO : BETTER ERROR HANDLING AND VALIDATION
			return c.Status(400).JSON("ID is not provided correctly")
		}

		c.Set("HX-Redirect", "/jobs"+"?"+"id="+strconv.Itoa(id))

		return c.SendStatus(200)
	})
}
