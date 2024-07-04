package profile

import (
	"encoding/json"
	"net/http"
)

func (c *ProfileController) GetOne(w http.ResponseWriter, r *http.Request) {
	body, err := parsePassport(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := c.ps.GetOne(r.Context(), body.PassportSerie, body.PassportNumber)
	if err != nil {
		//todo: добавить ошибку отсутствия записи
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
