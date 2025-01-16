package specification

import (
	"context"
	"github.com/google/uuid"
	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetByProjectIdCmd struct {
	ProjectId string
}

type GetByProjectIdCmdHandler struct {
	projectRepo interfaces.ProjectRepository
	specRepo    interfaces.SpecificationRepository
}

func NewGetByProjectIdCmdHandler(
	projectRepo interfaces.ProjectRepository,
	specRepo interfaces.SpecificationRepository) *GetByProjectIdCmdHandler {
	return &GetByProjectIdCmdHandler{projectRepo, specRepo}
}

func (h *GetByProjectIdCmdHandler) Handle(ctx context.Context, cmd *GetByProjectIdCmd) ([]*entity.Specification, error) {
	pid, err := uuid.Parse(cmd.ProjectId)
	if err != nil {
		return nil, err
	}

	project, err := h.projectRepo.FindById(ctx, pid)
	if err != nil {
		return nil, err
	}

	specifications, err := h.specRepo.FindByProjectId(ctx, project.Id)
	if err != nil {
		return nil, err
	}

	return specifications, nil
}
