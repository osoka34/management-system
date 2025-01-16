package project

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type DeleteProjectCmd struct {
	Id string
}

type DeleteProjectCmdHandler struct {
	repoProject interfaces.ProjectRepository
}

func NewDeleteProjectCmdHandler(
	repoProject interfaces.ProjectRepository,
) *DeleteProjectCmdHandler {
	return &DeleteProjectCmdHandler{
		repoProject: repoProject,
	}
}

func (h *DeleteProjectCmdHandler) Handle(
	ctx context.Context,
	cmd *DeleteProjectCmd,
) error {
	pid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return err
	}

	project, err := h.repoProject.FindById(ctx, pid)
	if err != nil {
		return err
	}

	project.Delete()

	if err := h.repoProject.UpdateProject(ctx, project); err != nil {
		return err
	}

	return nil
}
