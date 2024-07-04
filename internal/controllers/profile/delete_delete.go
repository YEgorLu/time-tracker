package profile

import "net/http"

func (c *ProfileController) Delete(w http.ResponseWriter, r *http.Request) {
	body, err := parsePassport(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.ps.Delete(r.Context(), body.PassportSerie, body.PassportNumber); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
