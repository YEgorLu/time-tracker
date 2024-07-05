package task

import (
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/service/timer"
	"github.com/YEgorLu/time-tracker/internal/util"
)

type taskController struct {
	timerS timer.TimerService
	log    logger.Logger
}

func NewController(timerS timer.TimerService, log logger.Logger) *taskController {
	return &taskController{timerS, log}
}

func (c taskController) RegisterRoute(router *http.ServeMux) {
	p := util.Rpm("tasks")
	router.HandleFunc(p(http.MethodPost, "timer", "start"), c.timerStart)
	router.HandleFunc(p(http.MethodGet, "timer", "{id}", "stop"), c.timerStop)
}
