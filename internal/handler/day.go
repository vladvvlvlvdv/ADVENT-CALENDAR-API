package handler

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/utils"
	"advent-calendar/pkg/validators"
	"time"

	"github.com/gofiber/fiber/v2"
)

var DayHandler = new(Handler)

// @Tags Days
// @Param Authorization header string true "Authorization"
// @Param request formData repository.DayDTO true "-"
// @Param attachments formData []file false " "
// @Success 200 {object} validators.GlobalErrorHandlerResp
// @Failure 400 {object} validators.GlobalErrorHandlerResp
// @Failure 500 {object} validators.GlobalErrorHandlerResp
// @Router /api/days [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	data := new(repository.DayDTO)

	files := c.Locals("files").(map[string][]utils.File)

	if err := c.BodyParser(data); err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(400, err.Error())
	}

	if errs := config.Validator.Validate(data); len(errs) > 0 && errs[0].Error {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return h.validateError(errs)
	}

	if err := repository.DayService.Create(*data, files["attachments"]); err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(500, "Ошибка при добавлении дня")
	}

	return c.JSON(validators.GlobalErrorHandlerResp{Success: true, Message: "День успешно добавлен"})
}

// @Tags Days
// @Success 200 {object} []repository.Day
// @Failrule 500 {object} validators.GlobalErrorHandlerResp
// @Router /api/days [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	setting, err := repository.SettingService.Get()
	if err != nil {
		return fiber.NewError(500, "Ошибка при получении настроек")
	}

	daysCount := utils.DaysInMonth(time.Now().Year(), time.Month(setting.Month))

	days, err := repository.DayService.GetAll(repository.Params{}, repository.Day{ID: uint(daysCount)})

	if err != nil {
		return fiber.NewError(500, "Ошибка при получении списка дней")
	}

	return c.JSON(days)
}
