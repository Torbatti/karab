package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/apis"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func Search(app *fiber.App) {
	app.Get("/search", func(c *fiber.Ctx) error {
		qrs := c.Queries()

		//
		log.Println("Queries")
		log.Println(qrs["name"])
		log.Println(qrs["city"])
		log.Println(qrs["speciality"])

		// Getting jobs from db
		var jobs []models.Job
		if qrs["name"] == "" {
			core.DataBase.Db.Limit(24).Find(&jobs)
			for _, job := range jobs {
				log.Println(job.Name)
			}
		} else {
			core.DataBase.Db.Limit(24).Find(&jobs, "name = ?", qrs["name"])
			for _, job := range jobs {
				log.Println(job.Name)
			}
		}

		var rjobs []apis.Job
		for _, job := range jobs {
			respJob := apis.ResponseJob(job)
			rjobs = append(rjobs, respJob)
		}

		return c.Render("pages/search", fiber.Map{
			"HeadTitle":   "جست و جو",
			"KeyWords":    "KeyWords",
			"Description": "Description",
			"Jobs":        rjobs,
		}, "layouts/main")
	})
}
