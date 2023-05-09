package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mBuergi86/GO-Microservices/src/handlers"
)

func Routes(app *fiber.App) {
	app.Get("/user", handlers.GetUsers)
	app.Get("/user/:id", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)
	app.Put("/user/:id", handlers.UpdateUser)
	app.Delete("/user/:id", handlers.DeleteUser)
}
