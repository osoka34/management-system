package user

import (
	"context"
	"service/internal/domain/entity"
	"service/internal/domain/interfaces"

	"github.com/google/uuid"
)

type CreateUserCmd struct {
	Login    string
	Password string
}

type CreateUserCmdHandler struct {
	repo interfaces.UserRepository
}


func NewCreateUserCmdHandler(repo interfaces.UserRepository) *CreateUserCmdHandler {
    return &CreateUserCmdHandler{repo: repo}
}


func (h *CreateUserCmdHandler) Handle(ctx context.Context, cmd *CreateUserCmd) (uuid.UUID, error) {
    usr := entity.NewUser(cmd.Login, cmd.Password)

    if err := h.repo.CreateUser(ctx, usr); err != nil {
        return uuid.Nil, err
    }

    return usr.Id.UUID(), nil
}


