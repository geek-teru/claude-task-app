package task

import (
	"fmt"

	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/domain/repository"
)

// Usecase はタスク操作のビジネスロジックを定義するインターフェース
type Usecase interface {
	Create(title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error)
	List() ([]*entity.Task, error)
	Get(id int64) (*entity.Task, error)
	Update(id int64, title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error)
	Delete(id int64) error
}

type usecase struct {
	repo repository.TaskRepository
}

// New は TaskUsecase の実装を返す
func New(repo repository.TaskRepository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
	task := &entity.Task{
		Title:       title,
		Description: description,
		Status:      status,
		UserID:      userID,
	}
	created, err := u.repo.Create(task)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to create task: %w", err)
	}
	return created, nil
}

func (u *usecase) List() ([]*entity.Task, error) {
	tasks, err := u.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to list tasks: %w", err)
	}
	return tasks, nil
}

func (u *usecase) Get(id int64) (*entity.Task, error) {
	task, err := u.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to get task: %w", err)
	}
	return task, nil
}

func (u *usecase) Update(id int64, title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
	// 既存タスクを取得して存在確認
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to find task: %w", err)
	}

	existing.Title = title
	existing.Description = description
	existing.Status = status
	existing.UserID = userID

	updated, err := u.repo.Update(existing)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to update task: %w", err)
	}
	return updated, nil
}

func (u *usecase) Delete(id int64) error {
	if err := u.repo.Delete(id); err != nil {
		return fmt.Errorf("usecase: failed to delete task: %w", err)
	}
	return nil
}
