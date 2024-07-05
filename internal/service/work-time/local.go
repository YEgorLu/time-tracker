package worktime

import (
	"context"
	"time"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/service/work-time/converters"
	"github.com/YEgorLu/time-tracker/internal/service/work-time/models"
	worktime "github.com/YEgorLu/time-tracker/internal/store/work-time"
	"github.com/google/uuid"
)

var _ WorkTimeService = &localWorkTimeService{}

type localWorkTimeService struct {
	store worktime.WorkTimeStore
	log   logger.Logger
}

func newLocalWorkTimeService(store worktime.WorkTimeStore, log logger.Logger) *localWorkTimeService {
	return &localWorkTimeService{
		store: store,
		log:   log,
	}
}

func (s *localWorkTimeService) Create(ctx context.Context, taskId uuid.UUID, ms int64) (models.WorkTime, error) {
	dbWorkTime, err := s.store.Create(ctx, taskId, ms)
	if err != nil {
		return models.WorkTime{}, nil
	}
	return converters.DbToSvWorkTime(dbWorkTime), nil
}

func (s *localWorkTimeService) GetByUser(ctx context.Context, userId uuid.UUID, start time.Time, end time.Time) ([]models.TaskWorktime, error) {
	dbTaskWorkTimes, err := s.store.GetByUser(ctx, userId, start, end)
	if err != nil {
		return []models.TaskWorktime{}, err
	}
	svTaskWorkTimes := converters.DbToSvTaskWorkTimeSlice(dbTaskWorkTimes)
	return svTaskWorkTimes, nil
}
