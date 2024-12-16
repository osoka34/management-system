package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetSpenRequirementsCmd struct {
	SpecificationId string
}

type GetSpenRequirementsCmdHandler struct {
	repoRequirement   interfaces.RequirementRepository
	repoSpecification interfaces.SpecificationRepository
}

func NewGetSpenRequirementsCmdHandler(
	repoRequirement interfaces.RequirementRepository,
	repoSpecification interfaces.SpecificationRepository,
) *GetSpenRequirementsCmdHandler {
	return &GetSpenRequirementsCmdHandler{
		repoRequirement:   repoRequirement,
		repoSpecification: repoSpecification,
	}
}

func (h *GetSpenRequirementsCmdHandler) Handle(
	ctx context.Context,
	cmd *GetSpenRequirementsCmd,
) ([]*entity.Requirement, error) {
	sid, err := uuid.Parse(cmd.SpecificationId)
	if err != nil {
		return nil, err
	}

	specification, err := h.repoSpecification.FindById(ctx, sid)
	if err != nil {
		return nil, err
	}

	requirements, err := h.repoRequirement.FindBySpecificationId(ctx, specification.Id)
	if err != nil {
		return nil, err
	}

	return requirements, nil
}
