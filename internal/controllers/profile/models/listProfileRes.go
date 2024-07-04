package models

import "github.com/YEgorLu/time-tracker/internal/service/profile/models"

type ListProfileRes struct {
	Data       []models.Profile `json:"data"`
	TotalCount int              `json:"totalCount"`
}
