package profile

import (
	"context"

	"github.com/YEgorLu/time-tracker/internal/logger"
	peopleinfo "github.com/YEgorLu/time-tracker/internal/service/peopleInfo"
	"github.com/YEgorLu/time-tracker/internal/service/profile/models"
	store "github.com/YEgorLu/time-tracker/internal/store/profile"
)

type ProfileService interface {
	GetOne(ctx context.Context, passportSerie, passportNumber string) (models.Profile, error)
	GetMany(ctx context.Context, page, size int, filter models.ProfileFilter) (profiles []models.Profile, count int, err error)
	Delete(ctx context.Context, passportSerie, passportNumber string) error
	Update(ctx context.Context, profile models.Profile) error
	Create(ctx context.Context, passportSerie, passportNumber string) (models.Profile, error)
}

func NewService(store store.ProfileStore, infoS peopleinfo.PeopleInfoService, log logger.Logger) *localProfileService {
	return newLocalProfileService(store, infoS, log)
}
