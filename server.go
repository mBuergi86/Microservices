package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mBuergi86/GO-Microservices/src/routes"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	routes.Routes(app)

	app.Static("/public", "./public/")

	log.Fatal(app.Listen(":8080"))
}
