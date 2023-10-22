package hx

import (
	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/apis"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func Base(app *fiber.App) {

	app.Get("/hx/index-search", func(c *fiber.Ctx) error {
		// Getting jobs from db
		var jobs []models.Job
		core.DataBase.Db.Limit(6).Find(&jobs)

		var rjobs []apis.Job
		for _, job := range jobs {
			respJob := apis.ResponseJob(job)
			rjobs = append(rjobs, respJob)
		}
		return c.Render("partials/index", fiber.Map{
			"Jobs": rjobs,
		})
	})
}
