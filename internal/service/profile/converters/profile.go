package converters

import (
	"github.com/YEgorLu/time-tracker/internal/service/profile/models"
	storeModels "github.com/YEgorLu/time-tracker/internal/store/profile/models"
)

func DbToSvProfile(m storeModels.Profile) models.Profile {
	return models.Profile{
		Name:           m.Name,
		Surname:        m.Surname,
		Patronymic:     m.Patronymic,
		Address:        m.Address,
		PassportNumber: m.PassportNumber,
		PassportSerie:  m.PassportSerie,
	}
}

func DbToSvProfileSlice(m []storeModels.Profile) []models.Profile {
	profiles := make([]models.Profile, len(m))
	for i, v := range m {
		profiles[i] = DbToSvProfile(v)
	}
	return profiles
}

func SvToDbProfile(m models.Profile) storeModels.Profile {
	return storeModels.Profile{
		Name:           m.Name,
		Surname:        m.Surname,
		Patronymic:     m.Patronymic,
		Address:        m.Address,
		PassportSerie:  m.PassportSerie,
		PassportNumber: m.PassportNumber,
	}
}
