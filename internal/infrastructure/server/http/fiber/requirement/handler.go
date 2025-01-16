package requirement

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"service/internal/app/command/requirement"
	"service/pkg/utils"
)

type RequirementHandler struct {
	createRequirementCmdHandler            *requirement.CreateRequirementCmdHandler
	updateRequirementCmdHandler            *requirement.UpdateRequirementCmdHandler
	deleteRequirementCmdHandler            *requirement.DeleteRequirementCmdHandler
	addInSpecRequirementCmdHandler         *requirement.AddInSpecificationCmdHandler
	getProjectRequirementsCmdHandler       *requirement.GetProjectRepuirementsCmdHandler
	getSpecificationRequirementsCmdHandler *requirement.GetSpecRequirementsCmdHandler
}

func NewRequirementHandler(
	createRequirementCmdHandler *requirement.CreateRequirementCmdHandler,
	updateRequirementCmdHandler *requirement.UpdateRequirementCmdHandler,
	deleteRequirementCmdHandler *requirement.DeleteRequirementCmdHandler,
	addInSpecRequirementCmdHandler *requirement.AddInSpecificationCmdHandler,
	getProjectRequirementsCmdHandler *requirement.GetProjectRepuirementsCmdHandler,
	getSpecificationRequirementsCmdHandler *requirement.GetSpecRequirementsCmdHandler,
) *RequirementHandler {
	return &RequirementHandler{
		createRequirementCmdHandler:            createRequirementCmdHandler,
		updateRequirementCmdHandler:            updateRequirementCmdHandler,
		deleteRequirementCmdHandler:            deleteRequirementCmdHandler,
		addInSpecRequirementCmdHandler:         addInSpecRequirementCmdHandler,
		getProjectRequirementsCmdHandler:       getProjectRequirementsCmdHandler,
		getSpecificationRequirementsCmdHandler: getSpecificationRequirementsCmdHandler,
	}
}

func (h *RequirementHandler) CreateRequirement(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = CreateRequirementRequest{}
	)

	l := zap.L().With(zap.String("method", "CreateRequirement"))
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

	uid, err := h.createRequirementCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).
		JSON(CreateRequirementResponse{Id: uid.String()})
}

func (h *RequirementHandler) UpdateRequirement(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = UpdateRequirementRequest{}
	)

	l := zap.L().With(zap.String("method", "UpdateRequirement"))
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

	err = h.updateRequirementCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).
		JSON(UpdateRequirementResponse{})
}

func (h *RequirementHandler) DeleteRequirement(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = DeleteRequirementRequest{}
	)

	l := zap.L().With(zap.String("method", "DeleteRequirement"))
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

	err = h.deleteRequirementCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).
		JSON(DeleteRequirementResponse{})
}

func (h *RequirementHandler) AddInSpecification(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = AddInSpecRequest{}
	)

	l := zap.L().With(zap.String("method", "AddInSpecification"))
	defer func() {
		if err != nil {
			l.Error("failed to add in specification", utils.SilentError(err))
		} else {
			l.Info("add in specification successfully")
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

	err = h.addInSpecRequirementCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(AddInSpecResponse{Ids: cmd.Ids})
}

func (h *RequirementHandler) GetProjectRequirements(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = GetProjectRequirementsRequest{}
	)

	l := zap.L().With(zap.String("method", "GetProjectRequirements"))
	defer func() {
		if err != nil {
			l.Error("failed to get project requirements", utils.SilentError(err))
		} else {
			l.Info("get project requirements successfully")
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

	requirements, err := h.getProjectRequirementsCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(NewGetProjectRequirementsResponse(requirements))
}

func (h *RequirementHandler) GetSpecificationRequirements(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = GetSpecificationRequirementsRequest{}
	)

	l := zap.L().With(zap.String("method", "GetSpecificationRequirements"))
	defer func() {
		if err != nil {
			l.Error("failed to get specification requirements", utils.SilentError(err))
		} else {
			l.Info("get specification requirements successfully")
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

	requirements, err := h.getSpecificationRequirementsCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(NewGetSpecificationRequirementsResponse(requirements))
}
