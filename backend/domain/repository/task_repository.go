package repository

import "github.com/nanch/claude-task-app/backend/domain/entity"

type TaskRepository interface {
	Create(task *entity.Task) (*entity.Task, error)
	FindAll() ([]*entity.Task, error)
	FindByID(id int64) (*entity.Task, error)
	Update(task *entity.Task) (*entity.Task, error)
	Delete(id int64) error
}
