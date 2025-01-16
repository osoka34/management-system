package specification

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"service/internal/domain/entity"
)

type SpecificationRepository struct {
	db *gorm.DB
}

func NewSpecificationRepository(db *gorm.DB) *SpecificationRepository {
	return &SpecificationRepository{db: db}
}

func (r *SpecificationRepository) CreateSpecification(
	ctx context.Context,
	specification *entity.Specification,
) error {
	daoSpec := FromEntity(specification)
	return r.db.WithContext(ctx).Create(daoSpec).Error
}

func (r *SpecificationRepository) FindById(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Specification, error) {
	var daoSpec SpecificationDAO
	if err := r.db.WithContext(ctx).First(&daoSpec, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return daoSpec.ToEntity(), nil
}

func (r *SpecificationRepository) FindByProjectId(
	ctx context.Context,
	projectId entity.ProjectId,
) ([]*entity.Specification, error) {
	var daoSpecs []SpecificationDAO
	if err := r.db.WithContext(ctx).Where("project_id = ? AND status = ?", projectId.UUID(), entity.StatusCreate).
		Find(&daoSpecs).Error; err != nil {
		return nil, err
	}

	var specifications []*entity.Specification
	for _, daoSpec := range daoSpecs {
		specifications = append(specifications, daoSpec.ToEntity())
	}

	return specifications, nil
}

func (r *SpecificationRepository) UpdateSpecification(
	ctx context.Context,
	specification *entity.Specification,
) error {
	daoSpec := FromEntity(specification)
	return r.db.WithContext(ctx).Save(daoSpec).Error
}
