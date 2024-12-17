package user

import "github.com/gofiber/fiber/v3"

func Map(router fiber.Router, h *UserHandler) {

    group := router.Group("/user")

	group.Post("/register", h.Register)
	group.Post("/login", h.Login)
}
