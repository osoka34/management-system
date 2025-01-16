package requirement

import "github.com/gofiber/fiber/v3"

func (h *RequirementHandler) Map(router fiber.Router) {
	group := router.Group("/requirement")

	group.Post("/create", h.CreateRequirement)
	group.Post("/update", h.UpdateRequirement)
	group.Post("/delete", h.DeleteRequirement)
	group.Post("/add_in_spec", h.AddInSpecification)
	group.Post("/get_project_requirements", h.GetProjectRequirements)
	group.Get("/get_spec_requirements", h.GetSpecificationRequirements)
}
