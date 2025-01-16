package project

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"service/internal/domain/entity"

	"service/internal/app/command/project"
	"service/pkg/utils"
)

type ProjectHandler struct {
	createCmdHandler         *project.CreateProjectCmdHandler
	updateCmdHandler         *project.UpdateProjectCmdHandler
	deleteCmdHandler         *project.DeleteProjectCmdHandler
	getAllProjectsCmdHandler *project.GetAllProjectsCmdHandler
	getProjectByIdCmdHandler *project.GetProjectByIdCmdHandler
}

func NewProjectHandler(
	createCmdHandler *project.CreateProjectCmdHandler,
	updateCmdHandler *project.UpdateProjectCmdHandler,
	deleteCmdHandler *project.DeleteProjectCmdHandler,
	getAllProjectsCmdHandler *project.GetAllProjectsCmdHandler,
	getProjectByIdCmdHandler *project.GetProjectByIdCmdHandler,
) *ProjectHandler {
	return &ProjectHandler{
		createCmdHandler:         createCmdHandler,
		updateCmdHandler:         updateCmdHandler,
		deleteCmdHandler:         deleteCmdHandler,
		getAllProjectsCmdHandler: getAllProjectsCmdHandler,
		getProjectByIdCmdHandler: getProjectByIdCmdHandler,
	}
}

func (h *ProjectHandler) CreateProject(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = CreateProjectRequest{}
	)

	l := zap.L().With(zap.String("method", "CreateProject"))
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

	uid, err := h.createCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).
		JSON(CreateProjectResponse{Id: uid.String()})
}

func (h *ProjectHandler) UpdateProject(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = UpdateProjectRequest{}
	)

	l := zap.L().With(zap.String("method", "UpdateProject"))
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

	err = h.updateCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).
		JSON(UpdateProjectResponse{in.Id})
}

func (h *ProjectHandler) DeleteProject(
	c fiber.Ctx,
) error {
	var (
		err error
		in  = DeleteProjectRequest{}
	)

	l := zap.L().With(zap.String("method", "DeleteProject"))
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

	err = h.deleteCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).
		JSON(DeleteProjectResponse{in.Id})
}

func (h *ProjectHandler) GetAllProjects(
	c fiber.Ctx,
) error {
	var (
		err error
	)

	l := zap.L().With(zap.String("method", "GetAllProjects"))
	defer func() {
		if err != nil {
			l.Error("failed to get all projects", utils.SilentError(err))
		} else {
			l.Info("get all projects successfully")
		}
	}()

	projects, err := h.getAllProjectsCmdHandler.Handle(c.UserContext(), &project.GetAllProjectsCmd{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(NewGetAllProjectsResponse(projects))
}

func (h *ProjectHandler) GetProjectById(
	c fiber.Ctx,
) error {

	var (
		err error
		in  = GetProjectByIdRequest{}
	)

	l := zap.L().With(zap.String("method", "GetProjectById"))
	defer func() {
		if err != nil {
			l.Error("failed to get project by id", utils.SilentError(err))
		} else {
			l.Info("get project by id successfully")
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

	var p *entity.Project
	p, err = h.getProjectByIdCmdHandler.Handle(c.UserContext(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.InternalErr(err))
	}

	return c.Status(fiber.StatusOK).JSON(NewGetProjectByIdResponse(p))

}
