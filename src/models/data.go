package models

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"sort"
)

type Data interface {
	GetUsers() []Users
	GetUserById(id int) Users
	CreateUser(Users) []Users
	UpdateUser(id int, user Users) []Users
	DeleteUser(id int) []Users
}

type data struct {
	userListe []Users
}

var (
	users    []Users
	filePath = "public/users.json"
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
)

func InitData() (Data, error) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := json.Unmarshal(file, &users); err != nil {
		return nil, fiber.ErrNotImplemented
	}

	return &data{users}, nil
}

func (d *data) GetUsers() []Users {
	return d.userListe
}

func (d *data) GetUserById(id int) Users {
	for _, user := range d.userListe {
		if user.Id == id {
			return user
		}
	}
	return Users{}
}

func (d *data) CreateUser(user Users) []Users {
	var (
		countID      = 1
		controlsUser Users
	)

	for _, item := range d.userListe {
		if item.Id != controlsUser.Id {
			countID++
		}
	}

	user.Id = countID

	d.userListe = append(d.userListe, user)

	sort.Slice(d.userListe, func(i, j int) bool {
		return d.userListe[i].Id < d.userListe[j].Id
	})

	newList, _ := json.MarshalIndent(d.userListe, "", " ")
	if err := os.WriteFile(filePath, newList, 0755); err != nil {
		log.Fatal(fiber.ErrExpectationFailed)
	}

	return d.userListe
}

func (d *data) UpdateUser(id int, newUser Users) []Users {
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

func (d *data) DeleteUser(id int) []Users {
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
