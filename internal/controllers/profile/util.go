package profile

import (
	"encoding/json"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
)

func parsePassport(r *http.Request) (models.GetProfileOneReq, error) {
	var body models.GetProfileOneReq
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}

func getNonEmptyStrings(s []string) []string {
	var notEmpty []string
	for _, v := range s {
		if len(s) > 0 {
			notEmpty = append(notEmpty, v)
		}
	}
	return notEmpty
}
