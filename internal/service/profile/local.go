package profile

import (
	"github.com/YEgorLu/time-tracker/internal/service/profile/models"
	store "github.com/YEgorLu/time-tracker/internal/store/profile"
)

var _ ProfileService = &localProfileService{}

type localProfileService struct {
	store store.ProfileStore
}

func newLocalProfileService(store store.ProfileStore) *localProfileService {
	return &localProfileService{store}
}

// Create implements ProfileService.
func (p *localProfileService) Create() (string, error) {
	panic("unimplemented")
}

// Delete implements ProfileService.
func (p *localProfileService) Delete(id string) error {
	panic("unimplemented")
}

// GetMany implements ProfileService.
func (p *localProfileService) GetMany() ([]models.Profile, error) {
	panic("unimplemented")
}

// GetOne implements ProfileService.
func (p *localProfileService) GetOne(id string) (models.Profile, error) {
	panic("unimplemented")
}

// Update implements ProfileService.
func (p *localProfileService) Update() error {
	panic("unimplemented")
}
