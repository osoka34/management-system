package project

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"service/internal/domain/entity"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(ctx context.Context, project *entity.Project) error {
	daoProject := FromEntity(project)
	return r.db.WithContext(ctx).Create(daoProject).Error
}

func (r *ProjectRepository) FindById(ctx context.Context, id uuid.UUID) (*entity.Project, error) {
	var daoProject ProjectDAO
	if err := r.db.WithContext(ctx).First(&daoProject, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Проект не найден
		}
		return nil, err // Ошибка базы данных
	}
	return daoProject.ToEntity(), nil
}

func (r *ProjectRepository) UpdateProject(ctx context.Context, project *entity.Project) error {
	daoProject := FromEntity(project)
	return r.db.WithContext(ctx).Save(daoProject).Error
}

func (r *ProjectRepository) FindByCreatorId(
	ctx context.Context,
	creatorId entity.UserId,
) ([]*entity.Project, error) {
	var daoProjects []ProjectDAO
	if err := r.db.WithContext(ctx).Where("creator_id = ?", creatorId.UUID()).Find(&daoProjects).Error; err != nil {
		return nil, err
	}

	var projects []*entity.Project
	for _, daoProject := range daoProjects {
		projects = append(projects, daoProject.ToEntity())
	}

	return projects, nil
}

func (r *ProjectRepository) AllCreatedProjects(ctx context.Context) ([]*entity.Project, error) {
	var daoProjects []ProjectDAO
	if err := r.db.WithContext(ctx).Where("status_id = ?", entity.StatusCreate).Find(&daoProjects).Error; err != nil {
		return nil, err
	}

	var projects []*entity.Project
	for _, daoProject := range daoProjects {
		projects = append(projects, daoProject.ToEntity())
	}

	return projects, nil
}
