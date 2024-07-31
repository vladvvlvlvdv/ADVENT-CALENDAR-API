package router

import (
	"advent-calendar/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func DayRouter(r fiber.Router) {
	users := r.Group("/day")
	users.Post("/", handler.DayHandler.Create)
}
