package models

import "github.com/google/uuid"

type Profile struct {
	Id             uuid.UUID
	Name           string
	Surname        string
	Patronymic     string
	Address        string
	PassportSerie  string
	PassportNumber string
}
