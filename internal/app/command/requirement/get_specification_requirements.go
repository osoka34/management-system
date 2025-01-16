package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetSpecRequirementsCmd struct {
	SpecificationId string
}

type GetSpecRequirementsCmdHandler struct {
	repoRequirement   interfaces.RequirementRepository
	repoSpecification interfaces.SpecificationRepository
}

func NewGetSpenRequirementsCmdHandler(
	repoRequirement interfaces.RequirementRepository,
	repoSpecification interfaces.SpecificationRepository,
) *GetSpecRequirementsCmdHandler {
	return &GetSpecRequirementsCmdHandler{
		repoRequirement:   repoRequirement,
		repoSpecification: repoSpecification,
	}
}

func (h *GetSpecRequirementsCmdHandler) Handle(
	ctx context.Context,
	cmd *GetSpecRequirementsCmd,
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
