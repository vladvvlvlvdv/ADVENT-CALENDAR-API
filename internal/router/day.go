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
	// days.Post("/", middleware.AdminMiddleware, utils.UploadFiles(dayUpload), handler.CreateDay)
	days.Put("/:id", middleware.AdminMiddleware, utils.UploadFiles(dayUpload), handler.UpdateDay)
	days.Get("/", handler.GetAllDays)
	days.Get("/admin", middleware.AdminMiddleware, handler.GetAllDaysForAdmin)
	days.Post("/:id/views", middleware.AuthMiddleware, handler.CreateDayView)

}
