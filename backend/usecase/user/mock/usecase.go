package mock

import "github.com/nanch/claude-task-app/backend/domain/entity"

// UserUsecase は user.Usecase のモック実装
type UserUsecase struct {
	CreateFn func(name, email string) (*entity.User, error)
	UpdateFn func(id int64, name, email string) (*entity.User, error)
}

func (m *UserUsecase) Create(name, email string) (*entity.User, error) {
	return m.CreateFn(name, email)
}

func (m *UserUsecase) Update(id int64, name, email string) (*entity.User, error) {
	return m.UpdateFn(id, name, email)
}
