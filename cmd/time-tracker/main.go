package main

import (
	_ "github.com/YEgorLu/time-tracker/api/rest"
	"github.com/YEgorLu/time-tracker/internal/config"
	"github.com/YEgorLu/time-tracker/internal/db"
	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/server"
)

// @title Time Tracker Swagger API
// @version 1.0
// @description Swagger API for Time Tracker.
// @termsOfService http://swagger.io/terms/

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /
func main() {
	l := logger.Get()
	serverConfig := server.ServerConfig{
		Port: config.App.Port,
		Log:  l,
	}
	defer db.CloseAll()
	err := server.
		NewServer(&serverConfig).
		Configure().
		WithSwagger().
		Run()
	if err != nil {
		l.Error(err)
	}
	l.Info("Program closed")
}
