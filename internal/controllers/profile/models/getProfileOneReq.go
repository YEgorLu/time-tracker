package models

type GetProfileOneReq struct {
	PassportSerie  string `json:"passportSerie"`
	PassportNumber string `json:"passportNumber"`
}
