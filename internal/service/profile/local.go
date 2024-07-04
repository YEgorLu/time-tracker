package profile

import (
	"context"

	"github.com/YEgorLu/time-tracker/internal/logger"
	peopleinfo "github.com/YEgorLu/time-tracker/internal/service/peopleInfo"
	"github.com/YEgorLu/time-tracker/internal/service/profile/converters"
	"github.com/YEgorLu/time-tracker/internal/service/profile/models"
	store "github.com/YEgorLu/time-tracker/internal/store/profile"
)

var _ ProfileService = &localProfileService{}

type localProfileService struct {
	store store.ProfileStore
	infoS peopleinfo.PeopleInfoService
	log   logger.Logger
}

func newLocalProfileService(store store.ProfileStore, infoS peopleinfo.PeopleInfoService, log logger.Logger) *localProfileService {
	return &localProfileService{store, infoS, log}
}

// Create implements ProfileService.
func (p *localProfileService) Create(ctx context.Context, passportSerie, passportNumber string) (models.Profile, error) {
	info, err := p.infoS.GetInfo(ctx, passportSerie, passportNumber)
	if err != nil {
		p.log.Error("error retrieving profile info", err)
	}
	profile := models.Profile{
		Name:           info.Name,
		Surname:        info.Surname,
		Patronymic:     info.Patronymic,
		Address:        info.Patronymic,
		PassportNumber: passportNumber,
		PassportSerie:  passportSerie,
	}
	dbProfile, err := p.store.Create(ctx, converters.SvToDbProfile(profile))
	if err != nil {
		return models.Profile{}, err
	}
	return converters.DbToSvProfile(dbProfile), nil
}

// Delete implements ProfileService.
func (p *localProfileService) Delete(ctx context.Context, passportSerie, passportNumber string) error {
	return p.store.Delete(ctx, passportSerie, passportNumber)
}

// GetMany implements ProfileService.
func (p *localProfileService) GetMany(ctx context.Context, page, size int, filter models.ProfileFilter) ([]models.Profile, int, error) {
	dbProfiles, count, err := p.store.GetMany(ctx, page, size, converters.SvToDbProfileFilter(filter))
	if err != nil {
		return []models.Profile{}, 0, err
	}
	return converters.DbToSvProfileSlice(dbProfiles), count, nil
}

// GetOne implements ProfileService.
func (p *localProfileService) GetOne(ctx context.Context, passportSerie, passportNumber string) (models.Profile, error) {
	dbProfile, err := p.store.GetOne(ctx, passportSerie, passportNumber)
	if err != nil {
		p.log.Error(err)
		return models.Profile{}, err
	}
	return converters.DbToSvProfile(dbProfile), err
}

// Update implements ProfileService.
func (p *localProfileService) Update(ctx context.Context, profile models.Profile) error {
	return p.store.Update(ctx, converters.SvToDbProfile(profile))
}
