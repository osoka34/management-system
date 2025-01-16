package project

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type UpdateProjectCmd struct {
	Id          string
	Title       string
	Description string
}

type UpdateProjectCmdHandler struct {
	repoProject interfaces.ProjectRepository
}

func NewUpdateProjectCmdHandler(
	repoProject interfaces.ProjectRepository,
) *UpdateProjectCmdHandler {
	return &UpdateProjectCmdHandler{
		repoProject: repoProject,
	}
}

func (h *UpdateProjectCmdHandler) Handle(
	ctx context.Context,
	cmd *UpdateProjectCmd,
) error {
	pid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return err
	}

	project, err := h.repoProject.FindById(ctx, pid)
	if err != nil {
		return err
	}

	project.UpdateTitle(cmd.Title)
	project.UpdateDescription(cmd.Description)

	if err := h.repoProject.UpdateProject(ctx, project); err != nil {
		return err
	}

	return nil
}
