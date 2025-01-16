package specification

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"service/internal/domain/entity"

	"service/internal/app/command/specification"
	"service/pkg/utils"
)

type SpecificationHandler struct {
	createSpecificationHandler *specification.CreateSpecificationCmdHandler
	updateSpecificationHandler *specification.UpdatedSpecificationHandler
	deleteSpecificationHandler *specification.DeleteSpecificationHandler
	getSpecByProjectIdHandler  *specification.GetByProjectIdCmdHandler
	getSpecByIdHandler         *specification.GetByIdCmdHandler
}

func NewSpecificationHandler(
	createSpecificationHandler *specification.CreateSpecificationCmdHandler,
	updateSpecificationHandler *specification.UpdatedSpecificationHandler,
	deleteSpecificationHandler *specification.DeleteSpecificationHandler,
	getSpecByProjectId *specification.GetByProjectIdCmdHandler,
	getSpecById *specification.GetByIdCmdHandler,
) *SpecificationHandler {
	return &SpecificationHandler{
		createSpecificationHandler: createSpecificationHandler,
		updateSpecificationHandler: updateSpecificationHandler,
		deleteSpecificationHandler: deleteSpecificationHandler,
		getSpecByProjectIdHandler:  getSpecByProjectId,
		getSpecByIdHandler:         getSpecById,
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

	return c.Status(fiber.StatusOK).JSON(CreateSpecificationResponse{Id: sid.String()})
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

func (h *SpecificationHandler) GetSpecByProjectId(c fiber.Ctx) error {
	var (
		err error
		in  GetSpecByProjectIdRequest
	)

	l := zap.L().With(zap.String("method", "GetSpecByProjectId"))
	defer func() {
		if err != nil {
			l.Error("failed to get", utils.SilentError(err))
		} else {
			l.Info("get successfully")
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

	var specs []*entity.Specification
	specs, err = h.getSpecByProjectIdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(NewGetSpecByProjectIdResponse(specs))
}

func (h *SpecificationHandler) GetSpecById(c fiber.Ctx) error {
	var (
		err error
		in  = GetSpecByIdRequest{}
	)

	l := zap.L().With(zap.String("method", "GetSpecById"))
	defer func() {
		if err != nil {
			l.Error("failed to get", utils.SilentError(err))
		} else {
			l.Info("get successfully")
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

	var spec *entity.Specification
	spec, err = h.getSpecByIdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(NewGetSpecByIdResponse(spec))

}
