package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetProjectRepuirementsCmd struct {
	ProjectId string
}

type GetProjectRepuirementsCmdHandler struct {
	repoRequirement interfaces.RequirementRepository
	repoProject     interfaces.ProjectRepository
}

func NewGetProjectRepuirementsCmdHandler(
	repoRequirement interfaces.RequirementRepository,
	repoProject interfaces.ProjectRepository,
) *GetProjectRepuirementsCmdHandler {
	return &GetProjectRepuirementsCmdHandler{
		repoRequirement: repoRequirement,
		repoProject:     repoProject,
	}
}

func (h *GetProjectRepuirementsCmdHandler) Handle(
	ctx context.Context,
	cmd *GetProjectRepuirementsCmd,
) ([]*entity.Requirement, error) {
	pid, err := uuid.Parse(cmd.ProjectId)
	if err != nil {
		return nil, err
	}

	project, err := h.repoProject.FindById(ctx, pid)
	if err != nil {
		return nil, err
	}

	requirements, err := h.repoRequirement.FindByProjectId(ctx, project.Id)
	if err != nil {
		return nil, err
	}

	return requirements, nil
}
