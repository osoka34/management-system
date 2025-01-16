package specification

import "github.com/gofiber/fiber/v3"

func (h *SpecificationHandler) Map(router fiber.Router) {
	group := router.Group("/specification")

	group.Post("/create", h.CreateSpecification)
	group.Post("/update", h.UpdateSpecification)
	group.Post("/delete", h.DeleteSpecification)
	group.Post("by_project_id", h.GetSpecByProjectId)
	group.Post("by_id", h.GetSpecById)
}
