package user

import (
    "service/internal/app/command/user"
    "github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	cmdGet *user.GetUserCmdHandler
	cmdReg *user.CreateUserCmdHandler
}

func NewUserHandler(
	cmdGet *user.GetUserCmdHandler,
	cmdReg *user.CreateUserCmdHandler,
) *UserHandler {
	return &UserHandler{cmdGet: cmdGet, cmdReg: cmdReg}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
}
