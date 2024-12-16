package project

import (
	"time"

	"github.com/google/uuid"

	"service/internal/domain/entity"
)

const tableName = "projects"

func (p *ProjectDAO) TableName() string {
    return tableName
}

type ProjectDAO struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey"` // UUID, соответствующий полю id в базе данных
	Title       string    `gorm:"type:varchar(255);not null"`
	StatusId    int64     `gorm:"type:bigint;not null"` // Статус проекта
	CreatorId   uuid.UUID `gorm:"type:uuid;not null"`   // UUID создателя проекта
	Description string    `gorm:"type:text;not null"`   // Описание проекта
	CreatedAt   time.Time `gorm:"not null"`             // Время создания
	UpdatedAt   time.Time `gorm:"not null"`             // Время обновления
}

func (p *ProjectDAO) ToEntity() *entity.Project {
	return &entity.Project{
		Id:          entity.ProjectId(p.Id),
		Title:       p.Title,
		StatusId:    entity.Status(p.StatusId),
		CreatorId:   entity.UserId(p.CreatorId),
		Description: p.Description,
		CreatedAt:   &p.CreatedAt,
		UpdatedAt:   &p.UpdatedAt,
	}
}

func FromEntity(project *entity.Project) *ProjectDAO {
	return &ProjectDAO{
		Id:          project.Id.UUID(),
		Title:       project.Title,
		StatusId:    int64(project.StatusId),
		CreatorId:   project.CreatorId.UUID(),
		Description: project.Description,
		CreatedAt:   *project.CreatedAt,
		UpdatedAt:   *project.UpdatedAt,
	}
}

