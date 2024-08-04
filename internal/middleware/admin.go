package middleware

import (
	"advent-calendar/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	token, err := utils.CheckBearerToken(c, "Authorization")
	if err != nil {
		return fiber.NewError(401, err.Error())
	}

	user, err := utils.VerifyToken(token)
	if err != nil {
		return fiber.NewError(403, "Нет доступа")
	}

	if user.Role != "admin" {
		return fiber.NewError(403, "Нет доступа")
	}

	c.Locals("user", user)

	return c.Next()
}
