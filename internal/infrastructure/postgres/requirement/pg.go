package requirement

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"service/internal/domain/entity"
)

type RequirementRepository struct {
	db *gorm.DB
}

func NewRequirementRepository(db *gorm.DB) *RequirementRepository {
	return &RequirementRepository{db: db}
}

func (r *RequirementRepository) CreateRequirement(
	ctx context.Context,
	requirement *entity.Requirement,
) error {
	daoRequirement := FromEntity(requirement)
	return r.db.WithContext(ctx).Create(daoRequirement).Error
}

func (r *RequirementRepository) FindById(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Requirement, error) {
	var daoRequirement RequirementDAO
	if err := r.db.WithContext(ctx).First(&daoRequirement, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return daoRequirement.ToEntity(), nil
}

func (r *RequirementRepository) FindByProjectId(
	ctx context.Context,
	projectId entity.ProjectId,
) ([]*entity.Requirement, error) {
	var daoRequirements []RequirementDAO
	if err := r.db.WithContext(ctx).Where("project_id = ?", projectId.UUID()).Find(&daoRequirements).Error; err != nil {
		return nil, err
	}

	var requirements []*entity.Requirement
	for _, daoReq := range daoRequirements {
		requirements = append(requirements, daoReq.ToEntity())
	}

	return requirements, nil
}

func (r *RequirementRepository) UpdateRequirement(
	ctx context.Context,
	requirement *entity.Requirement,
) error {
	daoRequirement := FromEntity(requirement)
	return r.db.WithContext(ctx).Save(daoRequirement).Error
}

func (r *RequirementRepository) FindBySpecificationId(
	ctx context.Context,
	specificationId entity.SpecificationId,
) ([]*entity.Requirement, error) {
	var daoRequirements []RequirementDAO
	if err := r.db.WithContext(ctx).Where("specification_id = ?", specificationId.UUID()).Find(&daoRequirements).Error; err != nil {
		return nil, err
	}

	var requirements []*entity.Requirement
	for _, daoReq := range daoRequirements {
		requirements = append(requirements, daoReq.ToEntity())
	}

	return requirements, nil
}

func (r *RequirementRepository) FindByExecutorId(
	ctx context.Context,
	executorId entity.UserId,
) ([]*entity.Requirement, error) {
	var daoRequirements []RequirementDAO
	if err := r.db.WithContext(ctx).Where("executor_id = ?", executorId.UUID()).Find(&daoRequirements).Error; err != nil {
		return nil, err
	}

	var requirements []*entity.Requirement
	for _, daoReq := range daoRequirements {
		requirements = append(requirements, daoReq.ToEntity())
	}

	return requirements, nil
}
