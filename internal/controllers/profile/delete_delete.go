package profile

import (
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/util"
)

// Delete Profile godoc
// @Summary Deletes given profile
// @Produce json
// @Param id path string true "Profile ID"
// @Success 200
// @Router /profile/{id} [delete]
func (c *ProfileController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := util.ParseUUID(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.ps.Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
