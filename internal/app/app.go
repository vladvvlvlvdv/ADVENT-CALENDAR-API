package app

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/router"
	"advent-calendar/pkg/validators"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type App struct {
	Server *fiber.App
}

func (a *App) Run() {
	a.Server = fiber.New(fiber.Config{
		AppName:      "ADVENT-CALENDAR-API",
		BodyLimit:    20 * 1024 * 1024 * 1024,
		ErrorHandler: validators.CustomErrorHandler,
	})

	a.Server.Use(logger.New(logger.Config{
		Format: "Time: ${time} [${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	a.Server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, RefreshToken",
	}))

	a.Server.Static("/api/public", "./public")

	if config.Config.MODE == "dev" {
		a.Server.Get("/swagger/*", swagger.HandlerDefault)
	}

	router.LoadRoutes(a.Server.Group("/api"))
	err := a.Server.Listen(fmt.Sprintf(":%s", config.Config.PORT))
	if err != nil {
		log.Fatal("Не удалось запустить сервер ", err)
	}
}
