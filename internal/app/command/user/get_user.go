package user

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"service/internal/domain/interfaces"
	"service/pkg/utils"
)

type GetUserCmd struct {
	Login    string
	Password string
}

type GetUserCmdHandler struct {
	repo interfaces.UserRepository
}

func NewGetUserCmdHandler(repo interfaces.UserRepository) *GetUserCmdHandler {
	return &GetUserCmdHandler{repo: repo}
}

func (h *GetUserCmdHandler) Handle(ctx context.Context, cmd *GetUserCmd) (uuid.UUID, error) {
	passHash := utils.HashSHA3(cmd.Password)
	usr, err := h.repo.FindByCreds(ctx, cmd.Login, passHash)
	if err != nil {
		return uuid.Nil, err
	}
	if usr == nil {
		return uuid.Nil, errors.New("user not found")
	}
	return usr.Id.UUID(), nil
}
