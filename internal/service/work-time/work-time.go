package worktime

import (
	"context"
	"time"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/service/work-time/models"
	worktimeStore "github.com/YEgorLu/time-tracker/internal/store/work-time"
	"github.com/google/uuid"
)

type WorkTimeService interface {
	Create(ctx context.Context, taskId uuid.UUID, ms int64) (models.WorkTime, error)
	GetByUser(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) ([]models.TaskWorktime, error)
}

func NewService(store worktimeStore.WorkTimeStore, log logger.Logger) WorkTimeService {
	return newLocalWorkTimeService(store, log)
}
