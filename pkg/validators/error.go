package validators

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	log.Printf("An error occurred: %v\n", err)

	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
	if err != nil {
		log.Printf("Failed to send error response: %v\n", err)
		return err
	}
	return nil
}
