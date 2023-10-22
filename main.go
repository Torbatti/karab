package main

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/torbatti/karab/apis"
	"github.com/torbatti/karab/core"
	"github.com/torbatti/karab/hx"
	"github.com/torbatti/karab/routes"
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

	// Auth
	routes.Auth(app)
	hx.Auth(app)
	hx.AuthSignup(app)
	hx.AuthLogin(app)

	// Applicant
	routes.Applicant(app)
	apis.ApplicantApis(app)

	// Company
	routes.Company(app)
	apis.CompanyApis(app)

	// Job
	routes.Job(app)
	apis.JobApis(app)

	// Search
	routes.Search(app)
	hx.Search(app)

	// Core
	routes.Base(app)
	hx.Base(app)

	// app.Stack() // Stack Trace
	app.Listen(":2233")
}
