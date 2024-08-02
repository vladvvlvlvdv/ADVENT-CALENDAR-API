package router

import (
	"advent-calendar/internal/handler"
	"advent-calendar/internal/middleware"
	"advent-calendar/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func DayRouter(r fiber.Router) {
	dayUpload := []utils.Upload{{FileKey: "attachments"}}

	days := r.Group("/days")
	days.Post("/", middleware.AdminMiddleware, utils.UploadFiles(dayUpload), handler.DayHandler.Create)
	days.Get("/", handler.DayHandler.GetAll)

}
