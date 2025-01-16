package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"service/internal/domain/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	daoUser := FromEntity(user)
	return r.db.WithContext(ctx).Create(daoUser).Error
}

func (r *UserRepository) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var daoUser UserDAO
	if err := r.db.WithContext(ctx).First(&daoUser, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Пользователь не найден
		}
		return nil, err // Ошибка базы данных
	}
	return daoUser.ToEntity(), nil
}

func (r *UserRepository) FindByCreds(
	ctx context.Context,
	login, hash string,
) (*entity.User, error) {
	var daoUser UserDAO
	if err := r.db.WithContext(ctx).First(&daoUser, "login = ? AND password_hash = ?", login, hash).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err // Пользователь не найден
		}
		return nil, err // Ошибка базы данных
	}
	return daoUser.ToEntity(), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	daoUser := FromEntity(user)
	return r.db.WithContext(ctx).Save(daoUser).Error
}


func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	var daoUsers []*UserDAO

	if err := r.db.WithContext(ctx).Find(&daoUsers).Error; err != nil {
		return nil, err
	}

	var users []*entity.User
	for _, daoUser := range daoUsers {
		users = append(users, daoUser.ToEntity())
	}

	return users, nil
}
