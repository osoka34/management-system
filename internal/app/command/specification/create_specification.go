package specification

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type CreateSpecificationCmd struct {
	ProjectId   string
	Title       string
	Description string
}

type CreateSpecificationCmdHandler struct {
	repoSpecification interfaces.SpecificationRepository
	repoProject       interfaces.ProjectRepository
}

func NewCreateSpecificationCmdHandler(
	repoSpecification interfaces.SpecificationRepository,
	repoProject interfaces.ProjectRepository,
) *CreateSpecificationCmdHandler {
	return &CreateSpecificationCmdHandler{
		repoSpecification: repoSpecification,
		repoProject:       repoProject,
	}
}

func (h *CreateSpecificationCmdHandler) Handle(
	ctx context.Context,
	cmd *CreateSpecificationCmd,
) (uuid.UUID, error) {
	pid, err := uuid.Parse(cmd.ProjectId)
	if err != nil {
		return uuid.Nil, err
	}

	project, err := h.repoProject.FindById(ctx, pid)
	if err != nil {
		return uuid.Nil, err
	}

	specification := entity.NewSpecification(
		cmd.Title,
		cmd.Description,
        project.Id,
	)

	if err := h.repoSpecification.CreateSpecification(ctx, specification); err != nil {
		return uuid.Nil, err
	}

	return specification.Id.UUID(), nil
}
