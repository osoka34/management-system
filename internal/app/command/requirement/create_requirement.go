package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type CreateRequirementCmd struct {
	ProjectId   string
	Title       string
	Description string
	ExecutorId  string
}

type CreateRequirementCmdHandler struct {
	repoRequirement interfaces.RequirementRepository
	repoProject     interfaces.ProjectRepository
	repoUser        interfaces.UserRepository
}

func NewCreateRequirementCmdHandler(
	repoRequirement interfaces.RequirementRepository,
	repoProject interfaces.ProjectRepository,
	repoUser interfaces.UserRepository,
) *CreateRequirementCmdHandler {
	return &CreateRequirementCmdHandler{
		repoRequirement: repoRequirement,
		repoProject:     repoProject,
		repoUser:        repoUser,
	}
}

func (h *CreateRequirementCmdHandler) Handle(
	ctx context.Context,
	cmd *CreateRequirementCmd,
) (uuid.UUID, error) {
	pid, err := uuid.Parse(cmd.ProjectId)
	if err != nil {
		return uuid.Nil, err
	}

	uid, err := uuid.Parse(cmd.ExecutorId)
	if err != nil {
		return uuid.Nil, err
	}

	project, err := h.repoProject.FindById(ctx, pid)
	if err != nil {
		return uuid.Nil, err
	}

	executor, err := h.repoUser.FindById(ctx, uid)
	if err != nil {
		return uuid.Nil, err
	}

	requirement := entity.NewRequirement(
		cmd.Title,
		cmd.Description,
		executor.Id,
		project.Id,
	)

	if err := h.repoRequirement.CreateRequirement(ctx, requirement); err != nil {
		return uuid.Nil, err
	}

	return requirement.Id.UUID(), nil
}
