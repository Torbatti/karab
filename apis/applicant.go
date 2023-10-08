package apis

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func ApplicantApis(app *fiber.App) {
	app.Post("/api/applicants/", CreateApplicant)
	app.Get("/api/applicants/", GetApplicants)
	app.Get("/api/applicants/:id", GetApplicant)
	app.Put("/api/applicants/:id", UpdateApplicant)
}

type Applicant struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	ApplicantName string `json:"applicant_name"`
}

func Sample(c *fiber.Ctx) {}

func CreateResponseApplicant(apl models.Applicant) Applicant {
	return Applicant{apl.ID, apl.ApplicantName}
}

func findApplicant(id int, apl *models.Applicant) error {
	core.DataBase.Db.Find(&apl, "id = ?", id)
	if apl.ID == 0 {
		return errors.New("Applicant Doesn't Exists")
	}
	return nil
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

func GetApplicants(c *fiber.Ctx) error {
	apls := []models.Applicant{}

	core.DataBase.Db.Find(&apls)
	respApls := []Applicant{}

	for _, apl := range apls {
		respApl := CreateResponseApplicant(apl)
		respApls = append(respApls, respApl)
	}

	return c.Status(200).JSON(respApls)
}

func GetApplicant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// TODO : BETTER ERROR HANDLING AND VALIDATION
		return c.Status(400).JSON("ID is not provided correctly")
	}

	var apl models.Applicant
	if err := findApplicant(id, &apl); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	respApl := CreateResponseApplicant(apl)

	return c.Status(200).JSON(respApl)
}

func UpdateApplicant(c *fiber.Ctx) error {
	// Params
	id, err := c.ParamsInt("id")
	if err != nil {
		// TODO : BETTER ERROR HANDLING AND VALIDATION
		return c.Status(400).JSON("ID is not provided correctly")
	}

	// Find Applicant
	var apl models.Applicant
	if err := findApplicant(id, &apl); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Update Applicant
	type UpdatedApplicant struct {
		ApplicantName string `json:"applicant_name"`
	}
	var updatedApplicant UpdatedApplicant
	if err := c.BodyParser(&updatedApplicant); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	apl.ApplicantName = updatedApplicant.ApplicantName

	// Save To db
	core.DataBase.Db.Save(&apl)
	respApl := CreateResponseApplicant(apl)
	return c.Status(200).JSON(respApl)
}
