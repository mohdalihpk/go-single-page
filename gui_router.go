package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func WebRoute() *fiber.App {
	app := fiber.New()
	// addFiltersMiddleWare(app)
	addStatic(app)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	// Create a new endpoint
	app.Get("/api", func(c *fiber.Ctx) error {
		// return c.SendString("Hello World!")
		return c.JSON(c.App().Stack())
	})

	return app
}

func addStatic(app *fiber.App) {
	app.Static("/", "./public")
}

// func addFiltersMiddleWare(app *fiber.App) {
// 	app.Use(func(c *fiber.Ctx) {
// 		// TODO: add middleware here
// 		c.Next()
// 	})
// }
