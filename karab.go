package main

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/torbatti/karab/apis"
	"github.com/torbatti/karab/core"
)

const Version = "00.00.01"

func main() {
	core.ConnectDb()
	engine := html.New("./views", ".html")
	engine.AddFunc(
		// add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Statics
	app.Static("/", "./public")

	// Apis
	apis.ApplicantApis(app)
	apis.CompanyApis(app)
	apis.JobApis(app)

	// Core
	core.Base(app)
	core.Auth(app)
	core.Search(app)
	core.Applicant(app)
	core.Company(app)
	core.Job(app)

	app.Stack() // Stack Trace
	app.Listen(":2233")
}
