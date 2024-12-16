package project

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type CreateProjectCmd struct {
	Title       string
	CreatorId   string
	Description string
}

type CreateProjectCmdHandler struct {
	repoProject interfaces.ProjectRepository
	repoUser    interfaces.UserRepository
}

func NewCreateProjectCmdHandler(
	repoProject interfaces.ProjectRepository,
	repoUser interfaces.UserRepository,
) *CreateProjectCmdHandler {
	return &CreateProjectCmdHandler{
		repoProject: repoProject,
		repoUser:    repoUser,
	}
}

func (h *CreateProjectCmdHandler) Handle(
	ctx context.Context,
	cmd *CreateProjectCmd,
) (uuid.UUID, error) {


    uid, err := uuid.Parse(cmd.CreatorId)
    if err != nil {
        return uuid.Nil, err
    }

	creator, err := h.repoUser.FindById(ctx, uid)
	if err != nil {
		return uuid.Nil, err
	}

	project := entity.NewProject(
		cmd.Title,
		creator.Id,
		cmd.Description,
	)

	if err := h.repoProject.CreateProject(ctx, project); err != nil {
		return uuid.Nil, err
	}

	return project.Id.UUID(), nil
}


