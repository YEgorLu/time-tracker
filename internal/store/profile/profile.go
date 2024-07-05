package profile

import (
	"context"
	"database/sql"

	"github.com/YEgorLu/time-tracker/internal/store/profile/models"
	"github.com/google/uuid"
)

type ProfileStore interface {
	GetMany(ctx context.Context, page, size int, filter models.ProfileFilter) (profiles []models.Profile, count int, err error)
	GetOne(ctx context.Context, id uuid.UUID) (models.Profile, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Create(context.Context, models.Profile) (models.Profile, error)
	Update(context.Context, models.Profile) error
}

func NewStore(conn *sql.DB) ProfileStore {
	return newPgProfileStore(conn)
}
