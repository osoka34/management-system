package project

import "github.com/gofiber/fiber/v3"

func (h *ProjectHandler) Map(router fiber.Router) {
	group := router.Group("/project")

	group.Post("/create", h.CreateProject)
	group.Post("/update", h.UpdateProject)
	group.Post("/delete", h.DeleteProject)
	group.Get("/list", h.GetAllProjects)
	group.Post("/by_id", h.GetProjectById)
}
