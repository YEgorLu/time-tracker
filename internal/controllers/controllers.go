package controllers

import (
	"net/http"
	"sync"

	"github.com/YEgorLu/time-tracker/internal/controllers/profile"
	"github.com/YEgorLu/time-tracker/internal/controllers/task"
	worktiimeController "github.com/YEgorLu/time-tracker/internal/controllers/work-time"
	"github.com/YEgorLu/time-tracker/internal/db"
	"github.com/YEgorLu/time-tracker/internal/logger"
	peopleinfo "github.com/YEgorLu/time-tracker/internal/service/peopleInfo"
	profileService "github.com/YEgorLu/time-tracker/internal/service/profile"
	"github.com/YEgorLu/time-tracker/internal/service/timer"
	worktimeService "github.com/YEgorLu/time-tracker/internal/service/work-time"
	profileStore "github.com/YEgorLu/time-tracker/internal/store/profile"
	timerStore "github.com/YEgorLu/time-tracker/internal/store/timer"
	worktimeStore "github.com/YEgorLu/time-tracker/internal/store/work-time"
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
	router.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {})
	return router, nil
}

func initControllers() {
	log := logger.Get()
	db := db.GetConnection(log)

	profileStore := profileStore.NewStore(db)
	timerStore := timerStore.NewStore(db, log)
	workTimeStore := worktimeStore.NewStore(db, log)

	peopleInfoService := peopleinfo.NewService(log)
	workTimeService := worktimeService.NewService(workTimeStore, log)
	timerService := timer.NewService(timerStore, workTimeService, log)

	profileService := profileService.NewService(profileStore, peopleInfoService, log)
	controllers = []Controller{
		profile.NewController(profileService, log),
		task.NewController(timerService, log),
		worktiimeController.NewController(workTimeService, log),
	}
}
