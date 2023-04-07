package main

import "log"



func main() {
	// app := fiber.New()

	app := WebRoute()
	// app.Static("/", "./public")
	// Start server on port 3000
	log.Fatal(app.Listen(":3000"))
}
