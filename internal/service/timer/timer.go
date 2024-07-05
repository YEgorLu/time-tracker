package timer

import (
	"context"

	"github.com/YEgorLu/time-tracker/internal/logger"
	worktime "github.com/YEgorLu/time-tracker/internal/service/work-time"
	"github.com/YEgorLu/time-tracker/internal/store/timer"
	"github.com/google/uuid"
)

type TimerService interface {
	Start(ctx context.Context, taskId, userId uuid.UUID) (uuid.UUID, error)
	Stop(ctx context.Context, timerId uuid.UUID) (elapsedMs int64, err error)
}

func NewService(store timer.TimerStore, workTimeS worktime.WorkTimeService, log logger.Logger) TimerService {
	return newLocalTimerService(store, workTimeS, log)
}
