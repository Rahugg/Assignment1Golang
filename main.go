package main

import (
	"github.com/gofiber/fiber/v2"

	"assignment1GO/database"
	"assignment1GO/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
