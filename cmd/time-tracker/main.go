package main

import (
	"fmt"

	_ "github.com/YEgorLu/time-tracker/api/rest"
	"github.com/YEgorLu/time-tracker/internal/config"
	"github.com/YEgorLu/time-tracker/internal/db"
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
	serverConfig := server.ServerConfig{
		Port: config.App.Port,
	}
	defer db.CloseAll()
	err := server.
		NewServer(&serverConfig).
		Configure().
		WithSwagger().
		Run()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Program closed")
}
