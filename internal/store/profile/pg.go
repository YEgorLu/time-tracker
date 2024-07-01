package profile

import (
	"database/sql"

	"github.com/YEgorLu/time-tracker/internal/store/profile/models"
)

var _ ProfileStore = &pgProfileStore{}

type pgProfileStore struct {
	conn *sql.DB
}

func newPgProfileStore(conn *sql.DB) *pgProfileStore {
	return &pgProfileStore{
		conn: conn,
	}
}

// Create implements ProfileStore.
func (p *pgProfileStore) Create(models.Profile) (string, error) {
	panic("unimplemented")
}

// Delete implements ProfileStore.
func (p *pgProfileStore) Delete(id string) error {
	panic("unimplemented")
}

// GetMany implements ProfileStore.
func (p *pgProfileStore) GetMany() ([]models.Profile, error) {
	panic("unimplemented")
}

// GetOne implements ProfileStore.
func (p *pgProfileStore) GetOne(id string) (models.Profile, error) {
	panic("unimplemented")
}

// Update implements ProfileStore.
func (p *pgProfileStore) Update(models.Profile) error {
	panic("unimplemented")
}
