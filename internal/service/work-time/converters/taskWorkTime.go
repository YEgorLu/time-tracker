package converters

import (
	"github.com/YEgorLu/time-tracker/internal/service/work-time/models"
	storeModels "github.com/YEgorLu/time-tracker/internal/store/work-time/models"
)

func DbToSvTaskWorkTime(m storeModels.TaskWorktime) models.TaskWorktime {
	return models.TaskWorktime{
		TaskId: m.TaskId,
		Ms:     m.Ms,
	}
}

func DbToSvTaskWorkTimeSlice(m []storeModels.TaskWorktime) []models.TaskWorktime {
	taskWorkTimes := make([]models.TaskWorktime, 0, len(m))
	for _, v := range m {
		taskWorkTimes = append(taskWorkTimes, DbToSvTaskWorkTime(v))
	}
	return taskWorkTimes
}
