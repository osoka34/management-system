package specification

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"service/internal/app/command/specification"
	"service/pkg/utils"
)

type SpecificationHandler struct {
	createSpecificationHandler *specification.CreateSpecificationCmdHandler
	updateSpecificationHandler *specification.UpdatedSpecificationHandler
	deleteSpecificationHandler *specification.DeleteSpecificationHandler
}

func NewSpecificationHandler(
	createSpecificationHandler *specification.CreateSpecificationCmdHandler,
	updateSpecificationHandler *specification.UpdatedSpecificationHandler,
	deleteSpecificationHandler *specification.DeleteSpecificationHandler,
) *SpecificationHandler {
	return &SpecificationHandler{
		createSpecificationHandler: createSpecificationHandler,
		updateSpecificationHandler: updateSpecificationHandler,
		deleteSpecificationHandler: deleteSpecificationHandler,
	}
}

func (h *SpecificationHandler) CreateSpecification(c fiber.Ctx) error {
	var (
		err error
		in  = CreateSpecificationRequest{}
	)

	l := zap.L().With(zap.String("method", "CreateSpecification"))
	defer func() {
		if err != nil {
			l.Error("failed to create", utils.SilentError(err))
		} else {
			l.Info("create successfully")
		}
	}()

	if err = c.Bind().JSON(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BadRequestErr(err))
	}

	if err = utils.ValidateStruct(&in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(utils.ValidationErr(err))
	}

	cmd := in.ToCmd()

	sid, err := h.createSpecificationHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusCreated).JSON(CreateSpecificationResponse{Id: sid.String()})
}

func (h *SpecificationHandler) UpdateSpecification(c fiber.Ctx) error {
	var (
		err error
		in  = UpdateSpecificationRequest{}
	)

	l := zap.L().With(zap.String("method", "UpdateSpecification"))
	defer func() {
		if err != nil {
			l.Error("failed to update", utils.SilentError(err))
		} else {
			l.Info("update successfully")
		}
	}()

	if err = c.Bind().JSON(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BadRequestErr(err))
	}

	if err = utils.ValidateStruct(&in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(utils.ValidationErr(err))
	}

	cmd := in.ToCmd()

	err = h.updateSpecificationHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(UpdateSpecificationResponse{Id: cmd.Id})
}

func (h *SpecificationHandler) DeleteSpecification(c fiber.Ctx) error {
	var (
		err error
		in  = DeleteSpecificationRequest{}
	)

	l := zap.L().With(zap.String("method", "DeleteSpecification"))
	defer func() {
		if err != nil {
			l.Error("failed to delete", utils.SilentError(err))
		} else {
			l.Info("delete successfully")
		}
	}()

	if err = c.Bind().JSON(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BadRequestErr(err))
	}

	if err = utils.ValidateStruct(&in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(utils.ValidationErr(err))
	}

	cmd := in.ToCmd()

	err = h.deleteSpecificationHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(DeleteSpecificationResponse{Id: cmd.Id})
}
