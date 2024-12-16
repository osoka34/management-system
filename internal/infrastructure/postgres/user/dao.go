package user

import (
	"time"

	"github.com/google/uuid"

	"service/internal/domain/entity"
)


const tableName = "users"

func (u *UserDAO) TableName() string {
    return tableName
}

type UserDAO struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey"` // UUID, соответствующий полю id в базе данных
	Login        string    `gorm:"type:varchar(255);not null;unique"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func (u *UserDAO) ToEntity() *entity.User {
	return &entity.User{
		Id:           entity.UserId(u.Id),
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
		CreatedAt:    &u.CreatedAt,
		UpdatedAt:    &u.UpdatedAt,
	}
}

func FromEntity(user *entity.User) *UserDAO {
	return &UserDAO{
		Id:           user.Id.UUID(),
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
		CreatedAt:    *user.CreatedAt,
		UpdatedAt:    *user.UpdatedAt,
	}
}
