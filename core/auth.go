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

	app.Post("/hx/login-btn", func(c *fiber.Ctx) error {
		l := new(AuthForm)
		if err := c.BodyParser(l); err != nil {
			return err
		}
		log.Println(l.UserName)
		log.Println(l.Password)

		// IMPLEMENT BASIC AUTH MIDDLEWARE
		// c.Set("HX-Push-Url", "/")
		c.Set("HX-Redirect", "/")

		log.Println(c.GetRespHeaders())

		// ONLY IF LOGGED IN RETURN 200
		return c.SendStatus(200)
	})

	app.Post("/hx/signup-btn", func(c *fiber.Ctx) error {
		l := new(AuthForm)
		if err := c.BodyParser(l); err != nil {
			return err
		}
		log.Println(l.UserName)
		log.Println(l.Password)
		// TODO: LOGIN AFTER SIGNUP
		// IMPLEMENT BASIC AUTH MIDDLEWARE
		// c.Set("HX-Push-Url", "/")
		c.Set("HX-Redirect", "/")

		log.Println(c.GetRespHeaders())

		// ONLY IF LOGGED IN RETURN 200
		return c.SendStatus(200)
	})

}
