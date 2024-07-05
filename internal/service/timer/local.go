package timer

import (
	"context"

	"github.com/YEgorLu/time-tracker/internal/logger"
	worktime "github.com/YEgorLu/time-tracker/internal/service/work-time"
	"github.com/YEgorLu/time-tracker/internal/store/timer"
	"github.com/google/uuid"
)

var _ TimerService = &localTimerService{}

type localTimerService struct {
	store     timer.TimerStore
	workTimeS worktime.WorkTimeService
	log       logger.Logger
}

func newLocalTimerService(store timer.TimerStore, workTimeService worktime.WorkTimeService, log logger.Logger) *localTimerService {
	return &localTimerService{
		store:     store,
		workTimeS: workTimeService,
		log:       log,
	}
}

func (s *localTimerService) Start(ctx context.Context, taskId, userId uuid.UUID) (uuid.UUID, error) {
	return s.store.Start(ctx, taskId, userId)
}

func (s *localTimerService) Stop(ctx context.Context, timerId uuid.UUID) (ms int64, err error) {
	timerRes, err := s.store.Stop(ctx, timerId)
	if err != nil {
		return 0, err
	}
	ms = timerRes.Elapsed().Milliseconds()
	if _, err := s.workTimeS.Create(ctx, timerRes.TaskId, ms); err != nil {
		return 0, err
	}
	return ms, nil
}
