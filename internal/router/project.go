package router

import (
	"advent-calendar/internal/handler"
	"advent-calendar/internal/middleware"
	"advent-calendar/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func ProjectRouter(r fiber.Router) {
	projUpload := []utils.Upload{{FileKey: "previews", FileType: "image", MaxCount: 1, Require: true}}
	projUploadUpd := []utils.Upload{{FileKey: "previews", FileType: "image", MaxCount: 1, Require: false}}

	projects := r.Group("/projects")
	projects.Post("/", middleware.AdminMiddleware, utils.UploadFiles(projUpload), handler.CreateProject)
	projects.Put("/:id", middleware.AdminMiddleware, utils.UploadFiles(projUploadUpd), handler.UpdateProject)
	projects.Delete("/:id", middleware.AdminMiddleware, handler.DeleteProject)
	projects.Get("/", handler.GetProjects)
}
