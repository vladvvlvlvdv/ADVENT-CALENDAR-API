package handler

import (
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

func CreateClick(c *fiber.Ctx) error {
	ip := c.IP()

	if err := repository.ClickService.Create(repository.Click{IP: ip}); err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "Просморт засчитан"})
}

func GetClicksCount(c *fiber.Ctx) error {
	count, err := repository.ClickService.Count()
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(fiber.Map{"count": count})
}
