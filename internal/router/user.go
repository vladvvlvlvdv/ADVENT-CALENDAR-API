package router

import (
	"advent-calendar/internal/handler"
	"advent-calendar/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(r fiber.Router) {
	users := r.Group("/users")
	users.Post("/login", handler.UserHandler.Login)
	users.Get("/check", middleware.AuthMiddleware, handler.UserHandler.Check)
}
