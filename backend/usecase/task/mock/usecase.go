package mock

import "github.com/nanch/claude-task-app/backend/domain/entity"

// TaskUsecase は task.Usecase のモック実装
type TaskUsecase struct {
	CreateFn func(title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error)
	ListFn   func() ([]*entity.Task, error)
	GetFn    func(id int64) (*entity.Task, error)
	UpdateFn func(id int64, title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error)
	DeleteFn func(id int64) error
}

func (m *TaskUsecase) Create(title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
	return m.CreateFn(title, description, status, userID)
}

func (m *TaskUsecase) List() ([]*entity.Task, error) {
	return m.ListFn()
}

func (m *TaskUsecase) Get(id int64) (*entity.Task, error) {
	return m.GetFn(id)
}

func (m *TaskUsecase) Update(id int64, title, description string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
	return m.UpdateFn(id, title, description, status, userID)
}

func (m *TaskUsecase) Delete(id int64) error {
	return m.DeleteFn(id)
}
