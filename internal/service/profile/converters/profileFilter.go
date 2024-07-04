package converters

import (
	"github.com/YEgorLu/time-tracker/internal/service/profile/models"
	storeModels "github.com/YEgorLu/time-tracker/internal/store/profile/models"
)

func SvToDbProfileFilter(m models.ProfileFilter) storeModels.ProfileFilter {
	return storeModels.ProfileFilter{
		Name:       m.Name,
		Surname:    m.Surname,
		Patronymic: m.Patronymic,
		Address:    m.Address,
	}
}
