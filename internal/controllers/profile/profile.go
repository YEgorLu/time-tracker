package profile

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YEgorLu/time-tracker/internal/logger"
)

type ProfileService interface{}

type ProfileController struct {
	ps  ProfileService
	log logger.Logger
}

type ProfileControllerProvider interface {
	GetProfileService() *ProfileService
	GetLogger() *log.Logger
}

func NewController(ps ProfileService, log logger.Logger) *ProfileController {
	return &ProfileController{
		ps,
		log,
	}
}

func (c *ProfileController) RegisterRoute(router *http.ServeMux) {
	basePath := "/profile"
	p := func(method, path string) string {
		if path == "" {
			return fmt.Sprintf("%s %s", method, basePath)
		}
		return fmt.Sprintf("%s %s/%s", method, basePath, path)
	}
	router.HandleFunc(p(http.MethodPost, ""), c.List)
	router.HandleFunc(p(http.MethodPost, ""), c.Create)
	router.HandleFunc(p(http.MethodDelete, "{id}"), c.Delete)
	router.HandleFunc(p(http.MethodPut, "{id}"), c.Update)
	router.HandleFunc(p(http.MethodGet, "{id}"), c.GetOne)
}

func (c *ProfileController) List(w http.ResponseWriter, r *http.Request) {
}

func (c *ProfileController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *ProfileController) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *ProfileController) Delete(w http.ResponseWriter, r *http.Request) {

}

func (c *ProfileController) GetOne(w http.ResponseWriter, r *http.Request) {

}
