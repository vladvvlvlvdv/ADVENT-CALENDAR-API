package handler

import (
	"advent-calendar/pkg/validators"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Handler func(c *fiber.Ctx)

func (h Handler) validateError(errs []validators.ErrorResponse) *fiber.Error {
	var errMsgs []string

	for _, err := range errs {
		var message string
		switch err.Tag {
		case "required":
			message = fmt.Sprintf("Поле '%s' обязательно для заполнения.", err.FailedField)
		case "email":
			message = fmt.Sprintf("Поле '%s' должно быть действительным адресом электронной почты.", err.FailedField)
		case "min":
			message = fmt.Sprintf("Поле '%s' должно содержать не менее %v символов.", err.FailedField, err.Param)
		case "max":
			message = fmt.Sprintf("Поле '%s' должно содержать не более %v символов.", err.FailedField, err.Param)
		case "len":
			message = fmt.Sprintf("Поле '%s' должно содержать ровно %v символов.", err.FailedField, err.Param)
		default:
			message = fmt.Sprintf("Поле '%s' имеет некорректное значение '%v'.", err.FailedField, err.Value)
		}
		errMsgs = append(errMsgs, message)
	}

	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, " "),
	}
}
