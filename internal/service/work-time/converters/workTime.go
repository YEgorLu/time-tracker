package converters

import (
	"github.com/YEgorLu/time-tracker/internal/service/work-time/models"
	storeModels "github.com/YEgorLu/time-tracker/internal/store/work-time/models"
)

func DbToSvWorkTime(m storeModels.WorkTime) models.WorkTime {
	return models.WorkTime{
		Id:        m.Id,
		TaskId:    m.TaskId,
		Ms:        m.Ms,
		Timestamp: m.Timestamp,
	}
}
