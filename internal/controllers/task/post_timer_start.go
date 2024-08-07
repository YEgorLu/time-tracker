package task

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/task/models"
)

type startRes models.TimerStartRes
type startReq models.TimerStartReq

// TimerStart godoc
// @Summary Starts timer on given task
// @Produce json
// @Param taskId body string true "Task ID"
// @Param userId body string true "User ID"
// @Success 200 {object} startRes
// @Router /timer/start [post]
func (c *taskController) timerStart(w http.ResponseWriter, r *http.Request) {
	var body models.TimerStartReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	timerId, err := c.timerS.Start(r.Context(), body.TaskId, body.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := models.TimerStartRes{TimerId: timerId}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
