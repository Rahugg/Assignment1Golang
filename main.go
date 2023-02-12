package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"assignment1GO/database"
	"assignment1GO/packages"
	"assignment1GO/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.Connect()

	//AllowCredentials is important because
	//front-end can get a cookie and send it back
	//if we want to authenticate using httponly cookies

	routes.Setup(app)

	html_read.HTML_read()
	html_read.Filter_words()
	html_read.GiveRating()

	app.Listen(":8000")
}
