package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/torbatti/karab/apis"
	"github.com/torbatti/karab/core"
)

const Version = "00.00.01"

func main() {
	engine := html.New("./views", ".html")
	core.ConnectDb()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Statics
	app.Static("/", "./public")

	// Apis
	apis.ApplicantApis(app)
	apis.CompanyApis(app)
	apis.JobApis(app)

	// Views
	core.BaseViews(app)
	core.ApplicantViews(app)
	core.CompanyViews(app)
	core.JobViews(app)

	app.Stack() // Stack Trace
	app.Listen(":2233")
}
