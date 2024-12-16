package requirement

import (
	"time"

	"github.com/google/uuid"

	"service/internal/domain/entity"
)

const tableName = "requirements"

func (r *RequirementDAO) TableName() string {
    return tableName
}

type RequirementDAO struct {
	Id              uuid.UUID  `gorm:"type:uuid;primaryKey"` // UUID, соответствующий полю id в базе данных
	Title           string     `gorm:"type:varchar(255);not null"`
	StatusId        int64      `gorm:"type:bigint;not null"` // Статус требования
	Description     string     `gorm:"type:text;not null"`   // Описание требования
	ExecutorId      uuid.UUID  `gorm:"type:uuid;not null"`   // UUID исполнителя
	ProjectId       uuid.UUID  `gorm:"type:uuid;not null"`   // UUID проекта
	SpecificationId *uuid.UUID `gorm:"type:uuid"`            // UUID спецификации (может быть NULL)
	CreatedAt       time.Time  `gorm:"not null"`
	UpdatedAt       time.Time  `gorm:"not null"`
}

func (r *RequirementDAO) ToEntity() *entity.Requirement {
	return &entity.Requirement{
		Id:              entity.RequirementId(r.Id),
		Title:           r.Title,
		StatusId:        entity.Status(r.StatusId),
		Description:     r.Description,
		ExecutorId:      entity.UserId(r.ExecutorId),
		ProjectId:       entity.ProjectId(r.ProjectId),
		SpecificationId: (*entity.SpecificationId)(r.SpecificationId),
		CreatedAt:       &r.CreatedAt,
		UpdatedAt:       &r.UpdatedAt,
	}
}

func FromEntity(requirement *entity.Requirement) *RequirementDAO {
	var specId *uuid.UUID
	if requirement.SpecificationId != nil {
		id := requirement.SpecificationId.UUID()
		specId = &id
	}

	return &RequirementDAO{
		Id:              requirement.Id.UUID(),
		Title:           requirement.Title,
		StatusId:        int64(requirement.StatusId),
		Description:     requirement.Description,
		ExecutorId:      requirement.ExecutorId.UUID(),
		ProjectId:       requirement.ProjectId.UUID(),
		SpecificationId: specId,
		CreatedAt:       *requirement.CreatedAt,
		UpdatedAt:       *requirement.UpdatedAt,
	}
}
