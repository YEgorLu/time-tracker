package profile

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile/converters"
	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
)

// Update Profile godoc
// @Summary Updates all Profile fields
// @Produce json
// @Param Id body string true "Profile ID"
// @Param PassportSerie body string true "Passport Serie"
// @Param PassportNumber body string true "Passport Number"
// @Param Name body string true "Name"
// @Param Surname body string true "Surname"
// @Param Patronymic body string false "Patronymic"
// @Param Address body string true "Address"
// @Success 200
// @Router /profile [put]
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
