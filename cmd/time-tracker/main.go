package main

import (
	"fmt"

	"github.com/YEgorLu/time-tracker/internal/config"
	"github.com/YEgorLu/time-tracker/internal/db"
	"github.com/YEgorLu/time-tracker/internal/server"
)

func main() {
	serverConfig := server.ServerConfig{
		Port: config.App.Port,
	}
	defer db.CloseAll()
	err := server.
		NewServer(&serverConfig).
		Configure().
		Run()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Program closed")
}
