package apis

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

func CompanyApis(app *fiber.App) {
	app.Post("/api/companies/", CreateCompany)
	app.Get("/api/companies/", GetCompanies)
	app.Get("/api/companies/:id", GetCompany)
	app.Put("/api/companies/:id", UpdateApplicant)
}

func ResponseCompany(cmp models.Company) Company {
	jobs := make([]Job, len(cmp.Jobs))
	for i, job := range cmp.Jobs {
		jobs[i] = Job{
			ID: job.ID,
		}
	}
	return Company{cmp.ID, cmp.Name, jobs}
}

func findCompanyById(id int, cmp *models.Company) error {
	core.DataBase.Db.Find(&cmp, "id = ?", id)
	if cmp.ID == 0 {
		return errors.New("Applicant Doesn't Exists")
	}
	return nil
}

func CreateCompany(c *fiber.Ctx) error {
	var cmp models.Company

	// TODO: VALIDATE
	if err := c.BodyParser(&cmp); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	core.DataBase.Db.Create(&cmp)

	resp := ResponseCompany(cmp)

	return c.Status(200).JSON(resp)
}

func GetCompanies(c *fiber.Ctx) error {
	cmps := []models.Company{}

	core.DataBase.Db.Find(&cmps)
	respCmps := []Company{}

	for _, cmp := range cmps {
		respCmp := ResponseCompany(cmp)
		respCmps = append(respCmps, respCmp)
	}

	return c.Status(200).JSON(respCmps)
}

func GetCompany(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// TODO : BETTER ERROR HANDLING AND VALIDATION
		return c.Status(400).JSON("ID is not provided correctly")
	}

	var cmp models.Company
	if err := findCompanyById(id, &cmp); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	resp := ResponseCompany(cmp)

	return c.Status(200).JSON(resp)
}
