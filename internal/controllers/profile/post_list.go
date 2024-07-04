package profile

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
	serviceModels "github.com/YEgorLu/time-tracker/internal/service/profile/models"
	"github.com/asaskevich/govalidator"
)

func (c *ProfileController) List(w http.ResponseWriter, r *http.Request) {
	var body models.ListProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if ok, err := govalidator.ValidateStruct(body); !ok || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}

	names := getNonEmptyStrings(body.Name)
	surnames := getNonEmptyStrings(body.Surname)
	patronymics := getNonEmptyStrings(body.Patronymic)
	addresses := getNonEmptyStrings(body.Patronymic)
	profiles, err := c.ps.GetMany(r.Context(), body.Page, body.Size, serviceModels.ProfileFilter{
		Name:       names,
		Surname:    surnames,
		Patronymic: patronymics,
		Address:    addresses,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(profiles); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
