package core

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type SearchForm struct {
	Name       string `form:"search"`
	City       string `form:"city"`
	Speciality string `form:"speciality"`
}

func Search(app *fiber.App) {

	app.Post("/clicked", func(c *fiber.Ctx) error {
		j := new(SearchForm)
		if err := c.BodyParser(j); err != nil {
			return err
		}
		log.Println(j.Name)
		log.Println(j.City)
		log.Println(j.Speciality)

		return c.Render("partials/clicked", fiber.Map{
			"Title": j.Name + " " + j.Speciality + " " + j.City,
		})
	})
}
