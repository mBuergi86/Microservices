package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/mBuergi86/GO-Microservices/src/models"
	"github.com/mBuergi86/GO-Microservices/src/repository"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func HelloWorld(c *fiber.Ctx) error {
	if err := c.Status(fiber.StatusOK).SendString("Hello World"); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500 Internal Server Error")
	}
	return nil
}

func ResultValue(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("value: " + c.Params("value"))
}

func GetUsers(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, err := repository.NewRepositoryUsers()
	if err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(data.GetUsers())
}

func GetUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := repository.NewRepositoryUsers()
	if err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.GetUserById(id))
}

func CreateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, err := repository.NewRepositoryUsers()
	if err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var user models.SUsers

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.CreateUser(user))
}

func UpdateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	var user models.SUsers
	data, err := repository.NewRepositoryUsers()
	if err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.UpdateUser(id, user))
}

func DeleteUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := repository.NewRepositoryUsers()
	if err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.DeleteUser(id))
}
