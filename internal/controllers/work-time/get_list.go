package worktime

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/YEgorLu/time-tracker/internal/controllers/work-time/converters"
	"github.com/google/uuid"
)

const (
	queryUserId         = "userId"
	queryTimestampStart = "timestampStart"
	queryTimestampEnd   = "timestampEnd"
)

var (
	queryParams = [...]string{queryUserId, queryTimestampStart, queryTimestampEnd}
)

func (c *workTimeController) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for _, v := range queryParams {
		if !query.Has(v) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	userIdQuery := query.Get(queryUserId)
	userId, err := uuid.Parse(userIdQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	timestampStartStr := query.Get(queryTimestampStart)
	timestampStart, err := time.Parse(time.RFC3339Nano, timestampStartStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	timestampEndStr := query.Get(queryTimestampEnd)
	timestampEnd, err := time.Parse(time.RFC3339Nano, timestampEndStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	svTaskWorkTimes, err := c.wrS.GetByUser(r.Context(), userId, timestampStart, timestampEnd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(converters.SbToCTaskWorkTimeSlice(svTaskWorkTimes)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
