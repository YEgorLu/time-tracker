package profile

import (
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/service/profile"
	"github.com/YEgorLu/time-tracker/internal/util"
)

type ProfileController struct {
	ps  profile.ProfileService
	log logger.Logger
}

func NewController(ps profile.ProfileService, log logger.Logger) *ProfileController {
	return &ProfileController{
		ps,
		log,
	}
}

func (c *ProfileController) RegisterRoute(router *http.ServeMux) {
	p := util.Rpm("/profile")
	c.log.Debug("path ", p(http.MethodPost))
	router.HandleFunc(p(http.MethodPost, "list"), c.List)
	router.HandleFunc(p(http.MethodPost), c.Create)
	router.HandleFunc(p(http.MethodDelete, "{id}"), c.Delete)
	router.HandleFunc(p(http.MethodPut), c.Update)
	router.HandleFunc(p(http.MethodGet, "{id}"), c.GetOne)
}
