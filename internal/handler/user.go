package handler

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

var UserHandler = new(Handler)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// @Tags Users
// @Param request formData repository.LoginDTO true "-"
// @Failure 401 {object} validators.GlobalErrorHandlerResp
// @Success 200 {object} Tokens
// @Router /api/users/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	data := new(repository.LoginDTO)

	if err := c.BodyParser(data); err != nil {
		return fiber.NewError(400, err.Error())
	}

	if errs := config.Validator.Validate(data); len(errs) > 0 && errs[0].Error {
		return h.validateError(errs)
	}

	user, err := repository.UserService.Get(repository.User{Email: data.Email})
	if err != nil {
		return fiber.NewError(401, "Неправильный логин")
	}

	if !utils.CheckPasswordHash(data.Password, user.Password) {
		return fiber.NewError(401, "Неправильный пароль")
	}

	jwt, err := utils.NewJWT(user.ID, user.Role)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	tokens := Tokens{AccessToken: jwt, RefreshToken: user.RefreshToken}

	return c.JSON(tokens)
}

// @Tags Users
// @Param Authorization header string true "Authorization"
// @Failure 401 {object} validators.GlobalErrorHandlerResp
// @Success 200 {object} repository.User
// @Router /api/users/check [get]
func (h *Handler) Check(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(utils.Claims)

	user, err := repository.UserService.Get(repository.User{ID: userClaims.ID})
	if err != nil {
		return fiber.NewError(401, "Ошибка при проверке авторизации")
	}

	return c.JSON(user)
}

// @Tags Users
// @Param RefreshToken header string true "RefreshToken"
// @Failure 500 {object} validators.GlobalErrorHandlerResp
// @Failure 401 {object} validators.GlobalErrorHandlerResp
// @Success 200 {object} Tokens
// @Router /api/users/refresh [patch]
func (h *Handler) Refresh(c *fiber.Ctx) error {
	refreshToken := c.Get("RefreshToken")

	user, err := repository.UserService.Get(repository.User{RefreshToken: refreshToken})
	if err != nil {
		return fiber.NewError(401, "Неправильный токен")
	}

	jwt, err := utils.NewJWT(user.ID, user.Role)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	newToken, err := utils.NewRefreshToken()
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	err = user.Update(repository.User{RefreshToken: newToken})
	if err != nil {
		return fiber.NewError(500, "Ошибка при обновлении токена")
	}

	tokens := Tokens{AccessToken: jwt, RefreshToken: user.RefreshToken}

	return c.JSON(tokens)
}
