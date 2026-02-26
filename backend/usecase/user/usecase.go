package user

import (
	"fmt"

	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/domain/repository"
)

// Usecase はユーザー操作のビジネスロジックを定義するインターフェース
type Usecase interface {
	Create(name, email string) (*entity.User, error)
	Update(id int64, name, email string) (*entity.User, error)
}

type usecase struct {
	repo repository.UserRepository
}

// New は UserUsecase の実装を返す
func New(repo repository.UserRepository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(name, email string) (*entity.User, error) {
	user := &entity.User{
		Name:  name,
		Email: email,
	}
	created, err := u.repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to create user: %w", err)
	}
	return created, nil
}

func (u *usecase) Update(id int64, name, email string) (*entity.User, error) {
	// 既存ユーザーを取得して存在確認
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to find user: %w", err)
	}

	existing.Name = name
	existing.Email = email

	updated, err := u.repo.Update(existing)
	if err != nil {
		return nil, fmt.Errorf("usecase: failed to update user: %w", err)
	}
	return updated, nil
}
