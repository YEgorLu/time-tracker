package models

type ListProfileRequest struct {
	Name       []string           `json:"name,omitempty"`
	Surname    []string           `json:"surname,omitempty"`
	Patronymic []string           `json:"patronymic,omitempty"`
	Address    []string           `json:"address,omitempty"`
	Passport   []GetProfileOneReq `json:"passport,omitempty"`

	Page int `json:"page" valid:"int,range(1)"`
	Size int `json:"size" valid:"int,range(1)"`
}
