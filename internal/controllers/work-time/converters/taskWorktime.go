package converters

import (
	"github.com/YEgorLu/time-tracker/internal/controllers/work-time/models"
	serviceModels "github.com/YEgorLu/time-tracker/internal/service/work-time/models"
)

func SvToCTaskWorkTime(m serviceModels.TaskWorktime) models.TaskWorktime {
	return models.TaskWorktime{
		TaskId: m.TaskId,
		Ms:     m.Ms,
	}
}

func SbToCTaskWorkTimeSlice(m []serviceModels.TaskWorktime) []models.TaskWorktime {
	taskWorkTimes := make([]models.TaskWorktime, 0, len(m))
	for _, v := range m {
		taskWorkTimes = append(taskWorkTimes, SvToCTaskWorkTime(v))
	}
	return taskWorkTimes
}
