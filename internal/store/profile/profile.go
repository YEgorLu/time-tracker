package profile

import (
	"database/sql"

	"github.com/YEgorLu/time-tracker/internal/store/profile/models"
)

type ProfileStore interface {
	GetMany() ([]models.Profile, error)
	GetOne(id string) (models.Profile, error)
	Delete(id string) error
	Create(models.Profile) (string, error)
	Update(models.Profile) error
}

func NewStore(conn *sql.DB) ProfileStore {
	return newPgProfileStore(conn)
}
