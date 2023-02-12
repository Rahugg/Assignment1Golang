package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(postgres.Open("postgres://postgres:12345@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the db")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
