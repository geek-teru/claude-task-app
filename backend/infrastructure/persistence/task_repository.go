package persistence

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/domain/repository"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *entity.Task) (*entity.Task, error) {
	model := toTaskModel(task)
	if err := r.db.Create(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	return toTaskEntity(&model), nil
}

func (r *taskRepository) FindAll() ([]*entity.Task, error) {
	var models []TaskModel
	zeroTime := time.Time{}
	if err := r.db.Where("deleted_at = ?", zeroTime).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find tasks: %w", err)
	}
	tasks := make([]*entity.Task, len(models))
	for i, m := range models {
		tasks[i] = toTaskEntity(&m)
	}
	return tasks, nil
}

func (r *taskRepository) FindByID(id int64) (*entity.Task, error) {
	var model TaskModel
	zeroTime := time.Time{}
	if err := r.db.Where("id = ? AND deleted_at = ?", id, zeroTime).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("task not found: %w", entity.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to find task: %w", err)
	}
	return toTaskEntity(&model), nil
}

func (r *taskRepository) Update(task *entity.Task) (*entity.Task, error) {
	model := toTaskModel(task)
	if err := r.db.Save(&model).Error; err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}
	return toTaskEntity(&model), nil
}

func (r *taskRepository) Delete(id int64) error {
	now := time.Now()
	if err := r.db.Model(&TaskModel{}).Where("id = ?", id).Update("deleted_at", now).Error; err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func toTaskModel(e *entity.Task) TaskModel {
	return TaskModel{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Status:      string(e.Status),
		UserID:      e.UserID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

func toTaskEntity(m *TaskModel) *entity.Task {
	return &entity.Task{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Status:      entity.TaskStatus(m.Status),
		UserID:      m.UserID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}
}
