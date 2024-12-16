package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type AddInSpecificationCmd struct {
	SpecificationId string
	Ids             []string
}

type AddInSpecificationCmdHandler struct {
	repoRequirement   interfaces.RequirementRepository
	repoSpecification interfaces.SpecificationRepository
}

func NewAddInSpecificationCmdHandler(
	repoRequirement interfaces.RequirementRepository,
	repoSpecification interfaces.SpecificationRepository,
) *AddInSpecificationCmdHandler {
	return &AddInSpecificationCmdHandler{
		repoRequirement:   repoRequirement,
		repoSpecification: repoSpecification,
	}
}

func (h *AddInSpecificationCmdHandler) Handle(
	ctx context.Context,
	cmd *AddInSpecificationCmd,
) error {
	sid, err := uuid.Parse(cmd.SpecificationId)
	if err != nil {
		return err
	}

	specification, err := h.repoSpecification.FindById(ctx, sid)
	if err != nil {
		return err
	}

	for _, id := range cmd.Ids {
		rid, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		requirement, err := h.repoRequirement.FindById(ctx, rid)
		if err != nil {
			return err
		}

        requirement.SetSpecification(specification.Id)

        if err := h.repoRequirement.UpdateRequirement(ctx, requirement); err != nil {
            return err
        }

	}

	return nil
}
