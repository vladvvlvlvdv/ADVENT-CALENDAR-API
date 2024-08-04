package router

import "github.com/gofiber/fiber/v2"

func LoadRoutes(r fiber.Router) {
	DayRouter(r)
	UserRouter(r)
	SettingRouter(r)
}
