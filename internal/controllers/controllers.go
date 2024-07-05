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
	// for i := 0; i < 10; i++ {
	// 	istr := strconv.Itoa(i)
	// 	serie := util.PadPrefix(istr, 4, "0")
	// 	number := "12345" + istr
	// 	dbprof := models.Profile{
	// 		Name:           "Name" + istr,
	// 		Surname:        "Surname" + istr,
	// 		Address:        "Address" + istr,
	// 		PassportSerie:  serie,
	// 		PassportNumber: number,
	// 	}
	// 	smth, err := profileStore.Create(context.Background(), dbprof)
	// 	if err != nil {
	// 		println(err.Error())
	// 	}
	// 	fmt.Println(smth)
	// }

	smth2, count, err2 := profileStore.GetMany(context.Background(), 1, 20, models.ProfileFilter{
		Name:    []string{"Name1", "Name2", "Name3"},
		Surname: []string{"Surname1", "Surname2", "Surname3"},
		Address: []string{"Address1", "Address2", "Address3"},
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
