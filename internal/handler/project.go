package handler

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/utils"
	"advent-calendar/pkg/validators"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

// @Tags Projects
// @Param Authorization header string true "Authorization"
// @Param request formData repository.ProjectDTO true "-"
// @Param preview formData file false " "
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 400 {object} validators.GlobalHandlerResp
// @Failure 500 {object} validators.GlobalHandlerResp
// @Router /api/projects [post]
func CreateProject(c *fiber.Ctx) error {
	data := new(repository.ProjectDTO)

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

	if err := repository.ProjectService.
		Create(repository.Project{
			Title:       data.Title,
			Description: data.Description,
			Preview:     files["previews"][0].Destination,
			Link:        data.Link,
		}); err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(500, "Ошибка при добавлении проекта")
	}

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "Проект добавлен"})
}

// @Tags Projects
// @Success 200 {object} []repository.Project
// @Failrule 500 {object} validators.GlobalHandlerResp
// @Router /api/projects [get]
func GetProjects(c *fiber.Ctx) error {
	projects, err := repository.ProjectService.GetAll(repository.Project{})

	if err != nil {
		return fiber.NewError(500, "Ошибка при получении списка проектов")
	}

	return c.JSON(projects)
}

// @Tags Projects
// @Param Authorization header string true "Authorization"
// @Param id path int true " "
// @Param request formData repository.ProjectDTO true "-"
// @Param preview formData file false " "
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 400 {object} validators.GlobalHandlerResp
// @Failure 500 {object} validators.GlobalHandlerResp
// @Router /api/projects/{id} [put]
func UpdateProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "Неверный ID проекта")
	}

	data := new(repository.ProjectDTO)

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

	project, err := repository.ProjectService.Get(repository.Project{ID: uint(id)})
	if err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(500, err.Error())
	}

	var previewToDel string
	if len(files) > 0 {
		if len(files["previews"]) > 0 {
			previewToDel = project.Preview
			project.Preview = files["previews"][0].Destination
		}
	}

	project.Title = data.Title
	project.Description = data.Description
	project.Link = data.Link

	if err := project.Update(); err != nil {
		if len(files) > 0 {
			if err := utils.DeleteFiles(files); err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return fiber.NewError(500, err.Error())
	}

	if previewToDel != "" {
		if err := os.Remove(fmt.Sprintf("./%s", previewToDel)); err != nil {
			return fiber.NewError(500, err.Error())
		}
	}

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "Проект изменен"})
}

// @Tags Projects
// @Param Authorization header string true "Authorization"
// @Param id path int true " "
// @Success 200 {object} validators.GlobalHandlerResp
// @Failure 400 {object} validators.GlobalHandlerResp
// @Failure 500 {object} validators.GlobalHandlerResp
// @Router /api/projects/{id} [delete]
func DeleteProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "Неверный ID проекта")
	}

	project, err := repository.ProjectService.Get(repository.Project{ID: uint(id)})
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	if err := project.Delete(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(validators.GlobalHandlerResp{Success: true, Message: "Проект удален"})
}
