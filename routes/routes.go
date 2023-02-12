package routes

import (
	"github.com/gofiber/fiber/v2"

	"assignment1GO/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
}