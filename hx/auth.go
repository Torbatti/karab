package hx

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/models"
)

type AuthForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func Auth(app *fiber.App) {

}

func AuthSignup(app *fiber.App) {

	app.Post("/hx/signup-btn", func(c *fiber.Ctx) error {
		l := new(AuthForm)
		if err := c.BodyParser(l); err != nil {
			return err
		}
		log.Println(l.UserName)
		log.Println(l.Password)

		// Create User
		var apl models.Applicant
		apl.ApplicantName = l.UserName
		core.DataBase.Db.Create(&apl)
		println(apl.ApplicantName)

		// TODO: LOGIN AFTER SIGNUP
		// IMPLEMENT BASIC AUTH MIDDLEWARE
		// c.Set("HX-Push-Url", "/")
		c.Set("HX-Redirect", "/")

		log.Println(c.GetRespHeaders())

		// ONLY IF LOGGED IN RETURN 200
		return c.SendStatus(200)
	})

}

func AuthLogin(app *fiber.App) {
	app.Post("/hx/login-btn", func(c *fiber.Ctx) error {
		l := new(AuthForm)
		if err := c.BodyParser(l); err != nil {
			return err
		}
		log.Println(l.UserName)
		log.Println(l.Password)

		// Find User
		var apl models.Applicant
		if err := core.DataBase.Db.Find(&apl, "applicant_name = ?", l.UserName).Error; err != nil {
			log.Println(err)
		}
		log.Println(apl.CreatedAt)

		// IMPLEMENT BASIC AUTH MIDDLEWARE
		// c.Set("HX-Push-Url", "/")
		c.Set("HX-Redirect", "/")

		log.Println(c.GetRespHeaders())

		// ONLY IF LOGGED IN RETURN 200
		return c.SendStatus(200)
	})
}
