package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mBuergi86/GO-Microservices/src/handlers"
)

func Routes(app *fiber.App) {
	app.Get("/user", handlers.ReadUsers)
	app.Get("/user/:id", handlers.ReadUserById)
	app.Post("/user", handlers.CreateUser)
	app.Put("/user/:id", handlers.UpdateUser)
	app.Delete("/user/:id", handlers.DeleteUser)
}

/*func GetUsers(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json; charset=utf-8")

	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": true,
			"msg":   "Ungültiger Content-Type. Bitte korrigieren Sie den Content-Type des Requests.",
		})
	}

	file, _ := os.ReadFile(filePath)
	if err := json.Unmarshal(file, &users); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error": true,
			"msg":   "Benutzer wurden nicht gefunden",
			"count": 0,
			"users": nil,
		})
	}

	return json.NewEncoder(c).Encode(users)
}*/

/*func GetUserById(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	file, _ := os.ReadFile(filePath)
	if err := json.Unmarshal(file, &users); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error": true,
			"msg":   "Benutzer wurden nicht gefunden",
			"count": 0,
			"users": nil,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	for _, item := range users {
		if item.Id == id {
			return json.NewEncoder(c).Encode(item)
		}
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}*/

/*func CreateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	file, _ := os.ReadFile(filePath)
	if err := json.Unmarshal(file, &users); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error": true,
			"msg":   "Benutzer wurden nicht gefunden",
			"count": 0,
			"users": nil,
		})
	}

	var user models.Users

	countID := 0

	for _, item := range users {
		if item.Id == countID {
			countID++
		}
	}

	user.Id = countID

	decoder := json.NewDecoder(bytes.NewReader(c.Body()))
	_ = decoder.Decode(&user)

	users = append(users, user)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})

	sortEncoder, _ := json.Marshal(users)

	encoder := json.NewEncoder(c)
	if err := encoder.Encode(sortEncoder); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, _ := json.MarshalIndent(users, "", " ")

	if err := os.WriteFile(filePath, data, 0755); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"msg":     "Benutzer erfolgreich erstellen",
		"data":    users,
	})
}*/

/*func UpdateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	file, _ := os.ReadFile(filePath)
	if err := json.Unmarshal(file, &users); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error": true,
			"msg":   "Benutzer wurden nicht gefunden",
			"count": 0,
			"users": nil,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	for index, item := range users {
		if item.Id == id {
			users = append(users[:index], users[index+1:]...)

			var user models.Users

			_ = json.NewDecoder(bytes.NewReader(c.Body())).Decode(&user)

			user.Id = len(users) + 1
			users = append(users, user)

			data, _ := json.MarshalIndent(users, "", " ")
			if err := os.WriteFile(filePath, data, 0755); err != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": true,
					"msg":   err.Error(),
				})
			}
			return json.NewEncoder(c).Encode(user)
		}
	}
	return json.NewEncoder(c).Encode(users)
}*/

/*func DeleteUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	file, _ := os.ReadFile(filePath)
	if err := json.Unmarshal(file, &users); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"error": true,
			"msg":   "Benutzer wurden nicht gefunden",
			"count": 0,
			"users": nil,
		})
	}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	for index, item := range users {
		if item.Id == id {
			users = append(users[:index], users[index+1:]...)
			data, _ := json.MarshalIndent(users, "", " ")
			if err := os.WriteFile(filePath, data, 0755); err != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": true,
					"msg":   err.Error(),
				})
			}
			break
		}
	}

	return json.NewEncoder(c).Encode(users)
}*/
