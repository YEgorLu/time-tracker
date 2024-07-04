package converters

import (
	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
	serviceModels "github.com/YEgorLu/time-tracker/internal/service/profile/models"
)

func SvToResProfile(m serviceModels.Profile) models.Profile {
	return models.Profile{
		PassportSerie:  m.PassportSerie,
		PassportNumber: m.PassportNumber,
		Name:           m.Name,
		Surname:        m.Surname,
		Patronymic:     m.Patronymic,
		Address:        m.Address,
	}
}

func ReqToSvProfile(m models.Profile) (serviceModels.Profile, error) {
	return serviceModels.Profile{
		Name:           m.Name,
		Surname:        m.Surname,
		Patronymic:     m.Patronymic,
		Address:        m.Address,
		PassportNumber: m.PassportNumber,
		PassportSerie:  m.PassportSerie,
	}, nil
}
