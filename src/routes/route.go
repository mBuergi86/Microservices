package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mBuergi86/GO-Microservices/src/handlers"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.HelloWorld)
	app.Get("/:value", handlers.ResultValue)
	app.Get("/users/user", handlers.GetUsers)
	app.Get("/users/user/:id", handlers.GetUser)
	app.Post("/users/user", handlers.CreateUser)
	app.Put("/users/user/:id", handlers.UpdateUser)
	app.Delete("/users/user/:id", handlers.DeleteUser)
}
