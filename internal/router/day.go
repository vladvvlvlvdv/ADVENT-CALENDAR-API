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
	days.Post("/", utils.UploadFiles(dayUpload), handler.CreateDay)
	days.Put("/:id", utils.UploadFiles(dayUpload), handler.UpdateDay)
	days.Get("/", handler.GetAllDays)
	days.Get("/admin", middleware.AdminMiddleware, handler.GetAllDaysForAdmin)

}
