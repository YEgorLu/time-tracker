package profile

import (
	"github.com/YEgorLu/time-tracker/internal/service/profile/models"
	store "github.com/YEgorLu/time-tracker/internal/store/profile"
)

type ProfileService interface {
	GetOne(id string) (models.Profile, error)
	GetMany() ([]models.Profile, error)
	Delete(id string) error
	Update() error
	Create() (string, error)
}

func NewService(store store.ProfileStore) *localProfileService {
	return newLocalProfileService(store)
}
