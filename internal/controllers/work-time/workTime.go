package worktime

import (
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/logger"
	worktimeService "github.com/YEgorLu/time-tracker/internal/service/work-time"
	"github.com/YEgorLu/time-tracker/internal/util"
)

type workTimeController struct {
	wrS worktimeService.WorkTimeService
	log logger.Logger
}

func NewController(wrS worktimeService.WorkTimeService, log logger.Logger) *workTimeController {
	return &workTimeController{
		wrS: wrS,
		log: log,
	}
}

func (c *workTimeController) RegisterRoute(router *http.ServeMux) {
	r := util.Rpm("work-time")
	router.HandleFunc(r(http.MethodGet, "list"), c.List)
}
