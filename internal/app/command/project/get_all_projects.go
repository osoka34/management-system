package project

import (
	"context"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetAllProjectsCmd struct{}

type GetAllProjectsCmdHandler struct {
	repoProject interfaces.ProjectRepository
}

func NewGetAllProjectsCmdHandler(
	repoProject interfaces.ProjectRepository,
) *GetAllProjectsCmdHandler {
	return &GetAllProjectsCmdHandler{repoProject: repoProject}
}

func (h *GetAllProjectsCmdHandler) Handle(
	ctx context.Context,
	cmd *GetAllProjectsCmd,
) ([]*entity.Project, error) {
	projects, err := h.repoProject.AllCreatedProjects(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
