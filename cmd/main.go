package main

import (
	"advent-calendar/internal/app"
	"advent-calendar/internal/config"
	"advent-calendar/internal/repository"
)

/* Swagger */
// @title Advent Calendar API docs
// @version 1.0

// @host localhost:9000
// @BasePath /

func main() {
	api := new(app.App)

	config.LoadConfig()
	repository.LoadDatabase()
	repository.AutoMigrate()
	repository.RenderDatabase()
	api.Run()
}
