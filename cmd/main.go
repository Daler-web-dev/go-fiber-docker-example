package main

import (
	"github.com/Daler-web-dev/go-fiber-docker-example/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Daler!")
	})

	app.Listen(":8080")
}
