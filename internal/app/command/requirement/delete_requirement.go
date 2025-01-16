package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type DeleteRequirementCmd struct {
	Id string
}

type DeleteRequirementCmdHandler struct {
	repoRequirement interfaces.RequirementRepository
}

func NewDeleteRequirementCmdHandler(
	repoRequirement interfaces.RequirementRepository,
) *DeleteRequirementCmdHandler {
	return &DeleteRequirementCmdHandler{
		repoRequirement: repoRequirement,
	}
}

func (h *DeleteRequirementCmdHandler) Handle(
	ctx context.Context,
	cmd *DeleteRequirementCmd,
) error {
	rid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return err
	}

	requirement, err := h.repoRequirement.FindById(ctx, rid)
	if err != nil {
		return err
	}

	requirement.Delete()

	if err := h.repoRequirement.UpdateRequirement(ctx, requirement); err != nil {
		return err
	}

	return nil
}
