package core

import "github.com/gofiber/fiber/v2"

func ApplicantViews(app *fiber.App) {
	app.Get("/applicant/:id")
	app.Get("/applicant/:name")
}
