package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/apis"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func Job(app *fiber.App) {
	app.Get("/jobs/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			// TODO : BETTER ERROR HANDLING AND VALIDATION
			return c.Status(400).JSON("ID is not provided correctly")
		}

		// Getting jobs from db
		var jobs []models.Job
		core.DataBase.Db.Limit(6).Find(&jobs)

		var rjobs []apis.Job
		for _, job := range jobs {
			respJob := apis.ResponseJob(job)
			rjobs = append(rjobs, respJob)
		}

		var job []models.Job
		core.DataBase.Db.Find(&job, "id =?", id)

		var rjob []apis.Job
		for _, job := range job {
			respJob := apis.ResponseJob(job)
			rjob = append(rjob, respJob)
		}

		return c.Render("pages/jobs", fiber.Map{
			"HeadTitle":   "جست و جو",
			"KeyWords":    "KeyWords",
			"Description": "Description",
			"Job":         rjob,
			"Jobs":        rjobs,
		}, "layouts/main")
	})
}
