package validators

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Param       interface{}
		Value       interface{}
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var Validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := Validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			elem.Param = err.Param()

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	log.Printf("An error occurred: %v\n", err)

	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = c.Status(code).JSON(GlobalErrorHandlerResp{Success: false, Message: e.Error()})
	if err != nil {
		log.Printf("Failed to send error response: %v\n", err)
		return err
	}
	return nil
}
