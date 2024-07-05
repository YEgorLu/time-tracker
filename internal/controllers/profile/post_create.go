package profile

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile/converters"
	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
	"github.com/asaskevich/govalidator"
)

type createRes models.Profile

// Create Profile godoc
// @Summary Creates profile using given passport data
// @Produce json
// @Param PassportNumber body string true "0000 000000"
// @Success 200
// @Router /profile [post]
func (c *ProfileController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body models.CreateProfileReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		if err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ok, err := govalidator.ValidateStruct(body); !ok || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	numbers := strings.Split(body.PassportNumber, " ")
	passportSerie, passportNumber := numbers[0], numbers[1]
	svProfile, err := c.ps.Create(ctx, passportSerie, passportNumber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(converters.SvToResProfile(svProfile)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
