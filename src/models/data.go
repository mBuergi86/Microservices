package models

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"sort"
)

type IData interface {
	GetUsers() []SUsers
	GetUserById(id int) SUsers
	CreateUser(SUsers) []SUsers
	UpdateUser(id int, user SUsers) []SUsers
	DeleteUser(id int) []SUsers
}

type Sdata struct {
	userListe []SUsers
}

var (
	users    []SUsers
	filePath = "public/users.json"
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
)

func InitData() (IData, error) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		return nil, fiber.ErrNotFound
	}

	if err := json.Unmarshal(file, &users); err != nil {
		return nil, fiber.ErrNotImplemented
	}

	return &Sdata{users}, nil
}

func (d *Sdata) GetUsers() []SUsers {
	return d.userListe
}

func (d *Sdata) GetUserById(id int) SUsers {
	for _, user := range d.userListe {
		if user.Id == id {
			return user
		}
	}
	return SUsers{}
}

func (d *Sdata) CreateUser(user SUsers) []SUsers {
	var (
		countID      = 1
		controlsUser SUsers
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

func (d *Sdata) UpdateUser(id int, newUser SUsers) []SUsers {
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

func (d *Sdata) DeleteUser(id int) []SUsers {
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
