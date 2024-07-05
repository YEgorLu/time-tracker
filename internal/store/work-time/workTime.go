package worktime

import (
	"context"
	"database/sql"
	"time"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/store/work-time/models"
	"github.com/google/uuid"
)

type WorkTimeStore interface {
	Create(ctx context.Context, taskId uuid.UUID, ms int64) (models.WorkTime, error)
	GetByUser(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) ([]models.TaskWorktime, error)
}

func NewStore(conn *sql.DB, log logger.Logger) WorkTimeStore {
	return newPgWorkTimeStore(conn, log)
}
