package middleware

import (
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token, err := utils.CheckBearerToken(c, "Authorization")
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

func WithoutAuthMiddleware(c *fiber.Ctx) error {
	token, _ := utils.CheckBearerToken(c, "Authorization")

	user, _ := utils.VerifyToken(token)

	dbUser, _ := repository.UserService.Get(repository.User{ID: user.ID})

	c.Locals("user", dbUser)

	return c.Next()
}
