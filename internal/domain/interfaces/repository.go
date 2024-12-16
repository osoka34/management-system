package interfaces

import (
	"context"

	"github.com/google/uuid"

	"service/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindById(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByCreds(ctx context.Context, login, hash string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
}

type SpecificationRepository interface {
	CreateSpecification(ctx context.Context, specification *entity.Specification) error
	FindById(ctx context.Context, id uuid.UUID) (*entity.Specification, error)
	FindByProjectId(
		ctx context.Context,
		projectId entity.ProjectId,
	) ([]*entity.Specification, error)
	UpdateSpecification(ctx context.Context, specification *entity.Specification) error
}

type RequirementRepository interface {
	CreateRequirement(ctx context.Context, requirement *entity.Requirement) error
	FindById(ctx context.Context, id uuid.UUID) (*entity.Requirement, error)
	FindByProjectId(ctx context.Context, projectId entity.ProjectId) ([]*entity.Requirement, error)
	UpdateRequirement(ctx context.Context, requirement *entity.Requirement) error
	FindBySpecificationId(
		ctx context.Context,
		specificationId entity.SpecificationId,
	) ([]*entity.Requirement, error)
	FindByExecutorId(
		ctx context.Context,
		executorId entity.UserId,
	) ([]*entity.Requirement, error)
}

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *entity.Project) error
	FindById(ctx context.Context, id uuid.UUID) (*entity.Project, error)
	UpdateProject(ctx context.Context, project *entity.Project) error
	FindByCreatorId(ctx context.Context, userId entity.UserId) ([]*entity.Project, error)
	AllCreatedProjects(ctx context.Context) ([]*entity.Project, error)
}
