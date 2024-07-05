package profile

import (
	"encoding/json"
	"net/http"
)

func (c *ProfileController) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := parseUUID(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profile, err := c.ps.GetOne(r.Context(), id)
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
