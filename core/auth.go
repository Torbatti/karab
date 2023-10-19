package core

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func Auth(app *fiber.App) {

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("pages/signup", fiber.Map{
			"HeadTitle":   "ثبت نام",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/login", fiber.Map{
			"HeadTitle":   "ورود",
			"KeyWords":    "KeyWords",
			"Description": "Description",
		}, "layouts/main")
	})

	app.Post("/dologin", func(c *fiber.Ctx) error {
		l := new(AuthForm)
		if err := c.BodyParser(l); err != nil {
			return err
		}
		log.Println(l.UserName)
		log.Println(l.Password)

		return c.SendStatus(400)
	})

	app.Post("/dosignup", func(c *fiber.Ctx) error {
		l := new(AuthForm)
		if err := c.BodyParser(l); err != nil {
			return err
		}
		log.Println(l.UserName)
		log.Println(l.Password)

		return c.SendStatus(400)
	})

}
