package handler

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

// @Tags Settings
// @Param Authorization header string true "Authorization"
// @Param request formData repository.SettingDTO true "-"
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 400 {object} validators.GlobalHandlerResp
// @Failure 500 {object} validators.GlobalHandlerResp
// @Router /api/settings [put]
func UpdateSettings(c *fiber.Ctx) error {
	data := new(repository.SettingDTO)

	if err := c.BodyParser(data); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if errs := config.Validator.Validate(data); len(errs) > 0 && errs[0].Error {
		return validators.ValidateError(errs)
	}

	if err := repository.SettingService.Update(repository.Setting{Month: data.Month, ShowAllDays: data.ShowAllDays}); err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "Настройки обновлены"})
}

// @Tags Settings
// @Param Authorization header string true "Authorization"
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 500 {object} repository.Setting
// @Router /api/settings [get]
func GetSettings(c *fiber.Ctx) error {
	settings, err := repository.SettingService.Get()
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(settings)
}
