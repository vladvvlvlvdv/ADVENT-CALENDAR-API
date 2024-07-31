package handler

import (
	"advent-calendar/internal/repository"

	"github.com/gofiber/fiber/v2"
)

var DayHandler = new(Handler)

// @Param request formData repository.DayDTO true "-"
// @Success 200 {object} repository.DayDTO
// @Router /api/day [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	data := new(repository.DayDTO)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	return c.JSON(data)
}
