package controllers

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile"
	"github.com/YEgorLu/time-tracker/internal/db"
	"github.com/YEgorLu/time-tracker/internal/logger"
	peopleinfo "github.com/YEgorLu/time-tracker/internal/service/peopleInfo"
	profileService "github.com/YEgorLu/time-tracker/internal/service/profile"
	profileStore "github.com/YEgorLu/time-tracker/internal/store/profile"
	"github.com/YEgorLu/time-tracker/internal/store/profile/models"
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
	db := db.GetConnection(log)
	profileStore := profileStore.NewStore(db)
	smth, err := profileStore.Create(context.Background(), models.Profile{
		Name:           "Name",
		Surname:        "Surname",
		Address:        "Address",
		PassportSerie:  "0000",
		PassportNumber: "123456",
	})
	if err != nil {
		println(err.Error())
	}
	fmt.Println(smth)
	smth2, count, err2 := profileStore.GetMany(context.Background(), 1, 2, models.ProfileFilter{
		Name:       []string{"name1", "name2", "name3"},
		Surname:    []string{"surname1", "surname2", "surname3"},
		Patronymic: []string{"patr1", "patr2", "patr3"},
		Address:    []string{"addr1", "addr2", "addr3"},
	})
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(smth2, count)
	peopleInfoService := peopleinfo.NewService(log)
	profileService := profileService.NewService(profileStore, peopleInfoService, log)
	controllers = []Controller{
		profile.NewController(profileService, log),
	}
}
