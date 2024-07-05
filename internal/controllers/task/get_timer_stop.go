package task

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/task/models"
	"github.com/YEgorLu/time-tracker/internal/util"
)

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
