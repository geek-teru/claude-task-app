package repository

import "github.com/nanch/claude-task-app/backend/domain/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByID(id int64) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}
