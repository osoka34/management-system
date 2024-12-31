package user

import (
	"service/internal/infrastructure/server/http/fiber/middleware"

	"github.com/gofiber/fiber/v3"
)

func Map(router fiber.Router, h *UserHandler) {

    group := router.Group("/user")

	group.Post("/register", h.Register)
	group.Post("/login", h.Login)


    authorization := group.Group("/auth", middleware.Authorization)

    authorization.Get("/refresh", h.RefreshToken)
}
