package middleware

import (
	"advent-calendar/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token, err := utils.GetBearerToken(c, "Authorization")
	if err != nil {
		return fiber.NewError(401, err.Error())
	}

	user, err := utils.VerifyToken(token)
	if err != nil {
		return fiber.NewError(403, "Нет доступа")
	}

	c.Locals("user", user)

	return c.Next()
}
