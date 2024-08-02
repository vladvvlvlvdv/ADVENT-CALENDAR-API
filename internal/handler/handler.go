package handler

import (
	"advent-calendar/pkg/validators"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Handler func(c *fiber.Ctx)

func (h Handler) validateError(errs []validators.ErrorResponse) *fiber.Error {
	errMsgs := make([]string, 0)

	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | должно быть '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}
	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, " и "),
	}
}
