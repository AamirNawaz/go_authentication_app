package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_authentication_app/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Users)
	app.Get("/user", controllers.GetUser)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
}
