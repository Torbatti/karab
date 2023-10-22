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

func Search(app *fiber.App) {

	app.Post("/hx/search-btn", func(c *fiber.Ctx) error {
		sf := new(SearchForm)
		if err := c.BodyParser(sf); err != nil {
			return err
		}

		log.Println("Form")
		log.Println(sf.Name)
		log.Println(sf.City)
		log.Println(sf.Speciality)

		c.Set("HX-Redirect", "/search"+"?"+"name="+sf.Name+"&city="+sf.City+"&speciality="+sf.Speciality)

		return c.Render("partials/clicked", fiber.Map{
			"Title": sf.Name + " " + sf.Speciality + " " + sf.City,
		})
	})
}
