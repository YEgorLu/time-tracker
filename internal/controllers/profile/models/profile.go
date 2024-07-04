package models

type Profile struct {
	PassportSerie  string `json:"passportSerie"`
	PassportNumber string `json:"passportNumber"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address"`
}
