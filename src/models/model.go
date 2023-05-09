package models

import "github.com/google/uuid"

type SUsers struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Vorname string    `json:"vorname"`
}
