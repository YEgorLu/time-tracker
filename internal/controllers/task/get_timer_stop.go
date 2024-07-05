package task

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/task/models"
	"github.com/YEgorLu/time-tracker/internal/util"
)

type s models.TimerStartRes

// TimerStop godoc
// @Summary Stops timer and creates worktime on task
// @Produce json
// @Param id path string true "Timer Id"
// @Success 200 {array} s
// @Router /tasks/timer/{id}/stop [get]
func (c *taskController) timerStop(w http.ResponseWriter, r *http.Request) {
	id, err := util.ParseUUID(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	passedTime, err := c.timerS.Stop(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := models.TimerStopRes{ElapsedMs: passedTime}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
