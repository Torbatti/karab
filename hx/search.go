package hx

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type SearchForm struct {
	Name       string `form:"search"`
	City       string `form:"city"`
	Speciality string `form:"speciality"`
}

func SearchHX(app *fiber.App) {

	app.Post("/hx/search-btn", func(c *fiber.Ctx) error {
		j := new(SearchForm)
		if err := c.BodyParser(j); err != nil {
			return err
		}
		log.Println(j.Name)
		log.Println(j.City)
		log.Println(j.Speciality)

		c.Set("HX-Push-Url", "/Test")
		log.Println(c.GetRespHeaders())

		return c.Render("partials/clicked", fiber.Map{
			"Title": j.Name + " " + j.Speciality + " " + j.City,
		})
	})
}
