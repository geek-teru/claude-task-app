package persistence

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/domain/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	model := toUserModel(user)
	if err := r.db.Create(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return toUserEntity(&model), nil
}

func (r *userRepository) FindByID(id int64) (*entity.User, error) {
	var model UserModel
	zeroTime := time.Time{}
	if err := r.db.Where("id = ? AND deleted_at = ?", id, zeroTime).First(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return toUserEntity(&model), nil
}

func (r *userRepository) Update(user *entity.User) (*entity.User, error) {
	model := toUserModel(user)
	if err := r.db.Save(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return toUserEntity(&model), nil
}

func toUserModel(e *entity.User) UserModel {
	return UserModel{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

func toUserEntity(m *UserModel) *entity.User {
	return &entity.User{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}
