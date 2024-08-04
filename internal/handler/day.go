package handler

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/utils"
	"advent-calendar/pkg/validators"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Tags Days
// @Param Authorization header string true "Authorization"
// @Param request formData repository.DayDTO true "-"
// @Param attachments formData []file false " "
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 400 {object} validators.GlobalHandlerResp
// @Failure 500 {object} validators.GlobalHandlerResp
// @Router /api/days [post]
func CreateDay(c *fiber.Ctx) error {
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
		return validators.ValidateError(errs)
	}

	if err := repository.DayService.Create(*data, files["attachments"]); err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(500, "Ошибка при добавлении дня")
	}

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "День успешно добавлен"})
}

// @Tags Days
// @Param timeZone query string false "Например Europe/Samara"
// @Success 200 {object} []repository.Day
// @Failrule 500 {object} validators.GlobalHandlerResp
// @Router /api/days [get]
func GetAllDays(c *fiber.Ctx) error {
	timeZone := c.Query("timeZone")

	setting, err := repository.SettingService.Get()
	if err != nil {
		return fiber.NewError(500, "Ошибка при получении настроек")
	}

	day := time.Now().Day()

	if timeZone != "" {
		day, err = utils.GetDayByTimeZone(timeZone)
		if err != nil {
			return fiber.NewError(500, "Некорректная временная зона")
		}
	}

	daysCount := utils.GetDaysCount(setting.Month, day, setting.ShowAllDays)

	days, err := repository.DayService.GetAll(repository.Params{}, repository.Day{ID: uint(daysCount)})

	if err != nil {
		return fiber.NewError(500, "Ошибка при получении списка дней")
	}

	return c.JSON(days)
}

// @Tags Days
// @Param Authorization header string true "Authorization"
// @Success 200 {object} []repository.Day
// @Failrule 500 {object} validators.GlobalHandlerResp
// @Router /api/days/admin [get]
func GetAllDaysForAdmin(c *fiber.Ctx) error {
	days, err := repository.DayService.GetAll(repository.Params{}, repository.Day{ID: 31})

	if err != nil {
		return fiber.NewError(500, "Ошибка при получении списка дней")
	}

	return c.JSON(days)
}

// @Tags Days
// @Param Authorization header string true "Authorization"
// @Param id path int true " "
// @Param request formData repository.DayUPD true "-"
// @Param attachments formData []file false " "
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 400 {object} validators.GlobalHandlerResp
// @Failure 500 {object} validators.GlobalHandlerResp
// @Router /api/days/{id} [put]
func UpdateDay(c *fiber.Ctx) error {
	files := c.Locals("files").(map[string][]utils.File)

	id, err := c.ParamsInt("id")
	if err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(400, "Некорректный ID")
	}

	day, err := repository.DayService.Get(repository.Day{ID: uint(id)})
	if err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(400, "День не найден")
	}

	data := new(repository.DayUPD)

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
		return validators.ValidateError(errs)
	}

	attachmentsToDel := make([]uint, 0)

	if len(data.AttachmentIds) > 0 {
		attachmentIdsMap := make(map[uint]bool)
		for _, attId := range data.AttachmentIds {
			attachmentIdsMap[attId] = true
		}

		for _, att := range day.Attachments {
			if !attachmentIdsMap[att.ID] {
				attachmentsToDel = append(attachmentsToDel, att.ID)
			}
		}
	} else {
		for _, att := range day.Attachments {
			attachmentsToDel = append(attachmentsToDel, att.ID)
		}
	}

	// if err := repository.DayService.Create(*data, files["attachments"]); err != nil {
	// 	if len(files) > 0 {
	// 		if err := utils.DeleteFiles(files); err != nil {
	// 			return fiber.NewError(500, err.Error())
	// 		}
	// 	}
	// 	return fiber.NewError(500, "Ошибка при добавлении дня")
	// }

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "День успешно обновлен"})
}
