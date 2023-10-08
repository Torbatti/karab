package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func ApplicantApis(app *fiber.App) {
	// app.Post("/api/applicant", CreateApplicant)
	// app.Get("/api/applicant/", GetApplicants)
	// app.Get("/api/applicant/:id", GetApplicant)
	// app.Put("/api/applicant/:id", CreateApplicant)
	// app.Delete("/api/applicant/:id", CreateApplicant)
}

type Applicant struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	ApplicantName string `json:"applicant_name"`
}

func Sample(c *fiber.Ctx) {}

func CreateResponseApplicant(apl models.Applicant) Applicant {
	return Applicant{apl.ID, apl.ApplicantName}
}

func CreateApplicant(c *fiber.Ctx) error {
	var apl models.Applicant

	// TODO: VALIDATE
	if err := c.BodyParser(&apl); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	core.DataBase.Db.Create(&apl)

	respApl := CreateResponseApplicant(apl)

	return c.Status(200).JSON(respApl)
}
