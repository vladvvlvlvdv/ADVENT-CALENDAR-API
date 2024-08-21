package router

import (
	"advent-calendar/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func ClickRouter(r fiber.Router) {
	clicks := r.Group("/clicks")

	clicks.Post("/", handler.CreateClick)
	clicks.Get("/", handler.GetClicksCount)
}
