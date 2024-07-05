package profile

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile/models"
	"github.com/google/uuid"
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

func parseUUID(idStr string) (uuid.UUID, error) {
	if idStr == "" {
		return uuid.Nil, errors.New("empty id")
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
