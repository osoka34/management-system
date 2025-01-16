package requirement

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type UpdateRequirementCmd struct {
	Id              string
	Title           string
	Description     string
	SpecificationId string
	ExecutorId      string
}

type UpdateRequirementCmdHandler struct {
	repoRequirement   interfaces.RequirementRepository
	repoSpecification interfaces.SpecificationRepository
	repoUser          interfaces.UserRepository
}

func NewUpdateRequirementCmdHandler(
	repoRequirement interfaces.RequirementRepository,
	repoSpecification interfaces.SpecificationRepository,
	repoUser interfaces.UserRepository,
) *UpdateRequirementCmdHandler {
	return &UpdateRequirementCmdHandler{
		repoRequirement:   repoRequirement,
		repoSpecification: repoSpecification,
		repoUser:          repoUser,
	}
}

func (h *UpdateRequirementCmdHandler) Handle(
	ctx context.Context,
	cmd *UpdateRequirementCmd,
) error {
	rid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return err
	}

	requirement, err := h.repoRequirement.FindById(ctx, rid)
	if err != nil {
		return err
	}

	if cmd.Title != "" {
		requirement.SetTitle(cmd.Title)
	}
	if cmd.Description != "" {
		requirement.SetDescription(cmd.Description)
	}

	if cmd.SpecificationId != "" {
		sid, err := uuid.Parse(cmd.SpecificationId)
		if err != nil {
			return err
		}

		specification, err := h.repoSpecification.FindById(ctx, sid)
		if err != nil {
			return err
		}

		requirement.SetSpecification(specification.Id)
	}

	if cmd.ExecutorId != "" {
		uid, err := uuid.Parse(cmd.ExecutorId)
		if err != nil {
			return err
		}

		executor, err := h.repoUser.FindById(ctx, uid)
		if err != nil {
			return err
		}

		requirement.SetExecutor(executor.Id)
	}

	if err := h.repoRequirement.UpdateRequirement(ctx, requirement); err != nil {
		return err
	}

	return nil
}
