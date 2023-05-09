package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/mBuergi86/GO-Microservices/src/models"
	"log"
	"os"
)

type IRepositoryUsers interface {
	GetUsers() []models.SUsers
	GetUserById(id uuid.UUID) models.SUsers
	CreateUser(models.SUsers) []models.SUsers
	UpdateUser(id uuid.UUID, user models.SUsers) []models.SUsers
	DeleteUser(id uuid.UUID) []models.SUsers
}

type SUsers struct {
	userListe []models.SUsers
}

var (
	users    []models.SUsers
	filePath = "public/users.json"
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
)

func NewRepositoryUsers() (*SUsers, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := json.Unmarshal(file, &users); err != nil {
		return nil, fiber.ErrNotImplemented
	}

	return &SUsers{users}, nil
}

func (d *SUsers) GetUsers() []models.SUsers {
	return d.userListe
}

func (d *SUsers) GetUserById(id uuid.UUID) models.SUsers {
	for _, user := range d.userListe {
		if user.Id == id {
			return user
		}
	}
	return models.SUsers{}
}

func (d *SUsers) CreateUser(user models.SUsers) []models.SUsers {
	countID := uuid.New()

	/*for _, item := range d.userListe {
		if item.Id != controlsUser.Id {
			countID++
		}
	}*/

	user.Id = countID

	d.userListe = append(d.userListe, user)

	/*sort.Slice(d.userListe, func(i, j int) bool {
		return d.userListe[i].Id < d.userListe[j].Id
	})*/

	newList, _ := json.MarshalIndent(d.userListe, "", " ")
	if err := os.WriteFile(filePath, newList, 0755); err != nil {
		log.Fatal(fiber.ErrExpectationFailed)
	}

	return d.userListe
}

func (d *SUsers) UpdateUser(id uuid.UUID, newUser models.SUsers) []models.SUsers {
	for i, user := range d.userListe {
		if user.Id == id {
			d.userListe[i] = newUser
			d.userListe[i].Id = id
			newList, _ := json.MarshalIndent(d.userListe, "", " ")
			if err := os.WriteFile(filePath, newList, 0755); err != nil {
				log.Fatal(fiber.ErrExpectationFailed)
			}
			break
		}
	}
	return d.userListe
}

func (d *SUsers) DeleteUser(id uuid.UUID) []models.SUsers {
	for i, user := range d.userListe {
		if user.Id == id {
			d.userListe = append(d.userListe[:i], d.userListe[i+1:]...)
			newList, _ := json.MarshalIndent(d.userListe, "", " ")
			if err := os.WriteFile(filePath, newList, 0755); err != nil {
				log.Fatal(fiber.ErrExpectationFailed)
			}
			break
		}
	}
	return d.userListe
}
