package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mBuergi86/GO-Microservices/src/routes"
	"log"
)

func HelloWorld(c *fiber.Ctx) error {
	if err := c.Status(fiber.StatusOK).SendString("Hello World"); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500 Internal Server Error")
	}
	return nil
}

func ResultValue(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("value: " + c.Params("value"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	routes.Routes(app)

	app.Get("/", HelloWorld)
	app.Get("/:value", ResultValue)
	app.Static("/public", "./public/")

	log.Fatal(app.Listen(":8080"))
}
