package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go_authentication_app/db"
	"go_authentication_app/routes"
)

func main() {
	//db connection
	db.Connect()

	app := fiber.New()

	//cors
	app.Use(cors.New())
	
	//routing
	routes.Setup(app)
	app.Listen(":3000")
}
