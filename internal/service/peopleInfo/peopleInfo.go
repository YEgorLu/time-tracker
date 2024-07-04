package peopleinfo

import (
	"context"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/service/peopleInfo/models"
)

type PeopleInfoService interface {
	GetInfo(ctx context.Context, passportSerie, passportNumber string) (models.Info, error)
}

func NewService(log logger.Logger) PeopleInfoService {
	return newExternalService(log)
}
