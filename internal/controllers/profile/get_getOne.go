package profile

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/util"
)

// Get Profile godoc
// @Summary Returns profile
// @Produce json
// @Param id path string true "Profile ID"
// @Success 200 {object} listRes
// @Router /profile/{id} [get]
func (c *ProfileController) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := util.ParseUUID(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := c.ps.GetOne(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
