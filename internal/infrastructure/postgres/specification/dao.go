package specification

import (
	"time"

	"github.com/google/uuid"

	"service/internal/domain/entity"
)


const tableName = "specifications"

func (s *SpecificationDAO) TableName() string {
    return tableName
}

type SpecificationDAO struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey"` // UUID, соответствующий полю id в базе данных
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	StatusId    int64     `gorm:"type:bigint;not null"` // Статус спецификации
	ProjectId   uuid.UUID `gorm:"type:uuid;not null"`   // UUID проекта
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

func (s *SpecificationDAO) ToEntity() *entity.Specification {
	return &entity.Specification{
		Id:          entity.SpecificationId(s.Id),
		Title:       s.Title,
		Description: s.Description,
		StatusId:    entity.Status(s.StatusId),
		ProjectId:   entity.ProjectId(s.ProjectId),
		CreatedAt:   &s.CreatedAt,
		UpdatedAt:   &s.UpdatedAt,
	}
}

func FromEntity(specification *entity.Specification) *SpecificationDAO {
	return &SpecificationDAO{
		Id:          specification.Id.UUID(),
		Title:       specification.Title,
		Description: specification.Description,
		StatusId:    int64(specification.StatusId),
		ProjectId:   specification.ProjectId.UUID(),
		CreatedAt:   *specification.CreatedAt,
		UpdatedAt:   *specification.UpdatedAt,
	}
}
