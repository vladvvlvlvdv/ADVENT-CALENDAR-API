package router

import (
	"advent-calendar/internal/handler"
	"advent-calendar/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SettingRouter(r fiber.Router) {
	settings := r.Group("/settings")
	settings.Put("/", middleware.AdminMiddleware, handler.UpdateSettings)
	settings.Get("/", middleware.AdminMiddleware, handler.GetSettings)
}
