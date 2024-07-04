package profile

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile/converters"
	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
)

func (c *ProfileController) Update(w http.ResponseWriter, r *http.Request) {
	var body models.Profile
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	svProfile, err := converters.ReqToSvProfile(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.ps.Update(r.Context(), svProfile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
