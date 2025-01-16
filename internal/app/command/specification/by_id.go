package specification

import (
	"context"
	"github.com/google/uuid"
	"service/internal/domain/entity"
	"service/internal/domain/interfaces"
)

type GetByIdCmd struct {
	Id string
}

type GetByIdCmdHandler struct {
	specRepo interfaces.SpecificationRepository
}

func NewGetByIdCmdHandler(specRepo interfaces.SpecificationRepository) *GetByIdCmdHandler {
	return &GetByIdCmdHandler{specRepo: specRepo}
}

func (h *GetByIdCmdHandler) Handle(ctx context.Context, cmd *GetByIdCmd) (*entity.Specification, error) {
	uid, err := uuid.Parse(cmd.Id)
	if err != nil {
		return nil, err
	}

	spec, err := h.specRepo.FindById(ctx, uid)
	if err != nil {
		return nil, err
	}

	return spec, nil
}
