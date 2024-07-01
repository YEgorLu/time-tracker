package controllers

import (
	"net/http"
	"sync"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile"
	"github.com/YEgorLu/time-tracker/internal/db"
	"github.com/YEgorLu/time-tracker/internal/logger"
	profileService "github.com/YEgorLu/time-tracker/internal/service/profile"
	profileStore "github.com/YEgorLu/time-tracker/internal/store/profile"
)

type Controller interface {
	RegisterRoute(*http.ServeMux)
}

var onceInit sync.Once
var controllers []Controller

func GetRoutes() (*http.ServeMux, error) {
	onceInit.Do(func() {
		initControllers()
	})
	router := http.NewServeMux()
	for _, controller := range controllers {
		controller.RegisterRoute(router)
	}
	return router, nil
}

func initControllers() {
	log := logger.Get()
	db, err := db.GetConnection()
	if err != nil {
		panic(err)
	}
	profileStore := profileStore.NewStore(db)
	profileService := profileService.NewService(profileStore)
	controllers = []Controller{
		profile.NewController(profileService, log),
	}
}
