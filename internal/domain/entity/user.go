package entity

import (
	"time"

	"github.com/google/uuid"

	"service/pkg/utils"
)

type UserId uuid.UUID

func (u UserId) UUID() uuid.UUID {
	return uuid.UUID(u)
}

func (u UserId) String() string {
	return uuid.UUID(u).String()
}

func NewUserId() UserId {
	return UserId(uuid.New())
}

type User struct {
	Id           UserId
	Login        string
	PasswordHash string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func NewUser(login, password string) *User {
	now := time.Now()
	return &User{
		Id:           NewUserId(),
		Login:        login,
		PasswordHash: utils.HashSHA3(password),
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
}
