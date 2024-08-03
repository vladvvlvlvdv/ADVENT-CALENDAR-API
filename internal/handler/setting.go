package handler

import "github.com/gofiber/fiber/v2"

var SettingHandler = new(Handler)

// @Tags Days
// @Param Authorization header string true "Authorization"
// @Param request formData repository.SettingDTO true "-"
// @Param attachments formData []file false " "
// @Success 200 {object} validators.GlobalErrorHandlerResp
// @Failure 400 {object} validators.GlobalErrorHandlerResp
// @Failure 500 {object} validators.GlobalErrorHandlerResp
// @Router /api/settings [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	return nil
}
