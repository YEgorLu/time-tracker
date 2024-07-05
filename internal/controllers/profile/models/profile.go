package models

import "github.com/google/uuid"

type Profile struct {
	Id             uuid.UUID `json:"id"`
	PassportSerie  string    `json:"passportSerie"`
	PassportNumber string    `json:"passportNumber"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Patronymic     string    `json:"patronymic,omitempty"`
	Address        string    `json:"address"`
}
