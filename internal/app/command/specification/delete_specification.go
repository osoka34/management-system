package specification

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
)

type DeleteSpecificationCmd struct {
	Id string
}

type DeleteSpecificationHandler struct {
	repoSpecification interfaces.SpecificationRepository
}

func NewDeleteSpecificationCmdHandler(
	repoSpecification interfaces.SpecificationRepository,
) *DeleteSpecificationHandler {
	return &DeleteSpecificationHandler{
		repoSpecification: repoSpecification,
	}
}

func (h *DeleteSpecificationHandler) Handle(
	ctx context.Context,
	cmd *DeleteSpecificationCmd,
) error {
	sid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return err
	}

	specification, err := h.repoSpecification.FindById(ctx, sid)
	if err != nil {
		return err
	}

	specification.Delete()

	return h.repoSpecification.UpdateSpecification(ctx, specification)
}
