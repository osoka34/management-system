package user

import (
	"context"
	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetAllCommand struct {
}

type GetAllUsersCmdHandler struct {
	userRepo interfaces.UserRepository
}

func NewGetAllUsersCmdHandler(userRepo interfaces.UserRepository) *GetAllUsersCmdHandler {
	return &GetAllUsersCmdHandler{
		userRepo: userRepo,
	}
}

func (c *GetAllUsersCmdHandler) Handle(ctx context.Context, _ *GetAllCommand) ([]*entity.User, error) {
	out, err := c.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return out, nil
}
