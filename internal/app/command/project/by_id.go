package project

import (
	"context"
	"github.com/google/uuid"
	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetProjectByIdCmd struct {
	Id string
}

type GetProjectByIdCmdHandler struct {
	projectRepo interfaces.ProjectRepository
}

func NewGetProjectByIdCmd(projectRepo interfaces.ProjectRepository) *GetProjectByIdCmdHandler {
	return &GetProjectByIdCmdHandler{projectRepo: projectRepo}
}

func (c *GetProjectByIdCmdHandler) Handle(ctx context.Context, cmd *GetProjectByIdCmd) (*entity.Project, error) {
	uid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return nil, err
	}
	project, err := c.projectRepo.FindById(ctx, uid)
	if err != nil {
		return nil, err
	}

	return project, nil
}
