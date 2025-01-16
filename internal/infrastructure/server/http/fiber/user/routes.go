package user

import (
	"github.com/gofiber/fiber/v3"

	"service/internal/infrastructure/server/http/fiber/middleware"
)

func (h *UserHandler) Map(router fiber.Router) {
	group := router.Group("/user")

	group.Post("/register", h.Register)
	group.Post("/login", h.Login)

	authorization := group.Group("/auth", middleware.Authorization)

	authorization.Get("/refresh", h.RefreshToken)
}
