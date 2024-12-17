package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"service/internal/app/command/user"
	"service/pkg/utils"
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

func (h *UserHandler) Register(c fiber.Ctx) error {
	var err error
	l := zap.L().With(zap.String("method", "Register"))
	defer func() {
		if err != nil {
			l.Error("failed to register", utils.SilentError(err))
		} else {
			l.Info("registered successfully")
		}
	}()

	var in RegisterRequest
	if err = c.Bind().JSON(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(utils.BadRequestErr(err))
	}

	if err = utils.ValidateStruct(&in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(utils.ValidationErr(err))
	}

	cmd := in.ToCmd()

	var uid uuid.UUID
	uid, err = h.cmdReg.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal server error",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(RegisterResponse{
			BearerToken: uid.String(),
		})
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	var err error
	l := zap.L().With(zap.String("method", "Login"))
	defer func() {
		if err != nil {
			l.Error("failed to login", utils.SilentError(err))
		} else {
			l.Info("logged in successfully")
		}
	}()

	var in Login
	if err = c.Bind().JSON(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(utils.BadRequestErr(err))
	}

	if err = utils.ValidateStruct(&in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(utils.ValidationErr(err))
	}

	cmd := in.ToCmd()

	var uid uuid.UUID
	uid, err = h.cmdGet.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal server error",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).
		JSON(RegisterResponse{
			BearerToken: uid.String(),
		})
}
