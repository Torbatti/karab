package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func JobApis(app *fiber.App) {
	app.Post("/api/Jobs/", CreateJob)
	app.Get("/api/Jobs/", GetJobs)
	app.Get("/api/Jobs/:id", GetCompany)
	app.Put("/api/Jobs/:id", UpdateApplicant)
}

func ResponseJob(job models.Job) Job {
	var cmp models.Company

	core.DataBase.Db.Find(&cmp, "id = ?", job.CompanyRefer)

	jobs := make([]Job, len(cmp.Jobs))
	for i, job := range cmp.Jobs {
		jobs[i] = Job{
			ID: job.ID,
		}
	}

	return Job{job.ID, job.Name, job.Description, job.City, job.Wage, job.Date, Company{cmp.ID, cmp.Name, jobs}}
}

// func findJobById(id int, job *models.Job) error {
// 	core.DataBase.Db.Find(&job, "id = ?", id)
// 	if job.ID == 0 {
// 		return errors.New("Applicant Doesn't Exists")
// 	}
// 	return nil
// }

func CreateJob(c *fiber.Ctx) error {
	var job models.Job

	// TODO: VALIDATE
	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	core.DataBase.Db.Create(&job)

	resp := ResponseJob(job)

	return c.Status(200).JSON(resp)
}

func GetJobs(c *fiber.Ctx) error {
	jobs := []models.Job{}

	core.DataBase.Db.Find(&jobs)
	respJobs := []Job{}

	for _, job := range jobs {
		respJob := ResponseJob(job)
		respJobs = append(respJobs, respJob)
	}

	return c.Status(200).JSON(respJobs)
}
