package specification

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type UpdateSpecificationCmd struct {
	Id          string
	Description string
	Title       string
}

type UpdatedSpecificationHandler struct {
	repoSpecification interfaces.SpecificationRepository
}

func NewUpdateSpecificationCmdHandler(
	repoSpecification interfaces.SpecificationRepository,
) *UpdatedSpecificationHandler {
	return &UpdatedSpecificationHandler{
		repoSpecification: repoSpecification,
	}
}

func (h *UpdatedSpecificationHandler) Handle(
	ctx context.Context,
	cmd *UpdateSpecificationCmd,
) error {
	sid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return err
	}

	specification, err := h.repoSpecification.FindById(ctx, sid)
	if err != nil {
		return err
	}

	if cmd.Title != "" {
		specification.SetTitle(cmd.Title)
	}
	if cmd.Description != "" {
		specification.SetDescription(cmd.Description)
	}

	return h.repoSpecification.UpdateSpecification(ctx, specification)
}
