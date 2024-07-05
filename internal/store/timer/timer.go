package timer

import (
	"context"
	"database/sql"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/store/timer/models"
	"github.com/google/uuid"
)

type TimerStore interface {
	Start(ctx context.Context, taskId, userId uuid.UUID) (uuid.UUID, error)
	Stop(ctx context.Context, timerId uuid.UUID) (models.TimerResult, error)
}

func NewStore(conn *sql.DB, log logger.Logger) TimerStore {
	return newPgStore(conn, log)
}
