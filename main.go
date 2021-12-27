package main

import (
	"github.com/gofiber/fiber/v2"
)

func anasayfa() string {
	return "Anasayfa"
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(anasayfa())
	})

	app.Listen(":3000")
}
