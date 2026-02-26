package task_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nanch/claude-task-app/backend/domain/entity"
	repomock "github.com/nanch/claude-task-app/backend/domain/repository/mock"
	"github.com/nanch/claude-task-app/backend/usecase/task"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		status      entity.TaskStatus
		userID      int64
		setup       func(ctrl *gomock.Controller) *repomock.MockTaskRepository
		want        *entity.Task
		wantErr     bool
	}{
		{
			name:        "正常系: タスクを作成できる",
			title:       "テストタスク",
			description: "説明",
			status:      entity.TaskStatusTodo,
			userID:      1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().Create(gomock.Any()).DoAndReturn(func(t *entity.Task) (*entity.Task, error) {
					t.ID = 1
					return t, nil
				})
				return m
			},
			want: &entity.Task{
				ID:          1,
				Title:       "テストタスク",
				Description: "説明",
				Status:      entity.TaskStatusTodo,
				UserID:      1,
			},
		},
		{
			name:        "異常系: リポジトリエラー",
			title:       "テストタスク",
			description: "説明",
			status:      entity.TaskStatusTodo,
			userID:      1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().Create(gomock.Any()).Return(nil, errors.New("db error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := task.New(tt.setup(ctrl))
			got, err := uc.Create(tt.title, tt.description, tt.status, tt.userID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{}, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("Create() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(ctrl *gomock.Controller) *repomock.MockTaskRepository
		want    []*entity.Task
		wantErr bool
	}{
		{
			name: "正常系: タスク一覧を取得できる",
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindAll().Return([]*entity.Task{
					{ID: 1, Title: "タスク1", Status: entity.TaskStatusTodo, UserID: 1},
					{ID: 2, Title: "タスク2", Status: entity.TaskStatusDone, UserID: 1},
				}, nil)
				return m
			},
			want: []*entity.Task{
				{ID: 1, Title: "タスク1", Status: entity.TaskStatusTodo, UserID: 1},
				{ID: 2, Title: "タスク2", Status: entity.TaskStatusDone, UserID: 1},
			},
		},
		{
			name: "正常系: タスクが0件でも空スライスを返す",
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindAll().Return([]*entity.Task{}, nil)
				return m
			},
			want: []*entity.Task{},
		},
		{
			name: "異常系: リポジトリエラー",
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindAll().Return(nil, errors.New("db error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := task.New(tt.setup(ctrl))
			got, err := uc.List()
			if (err != nil) != tt.wantErr {
				t.Fatalf("List() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{}, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("List() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		id      int64
		setup   func(ctrl *gomock.Controller) *repomock.MockTaskRepository
		want    *entity.Task
		wantErr bool
	}{
		{
			name: "正常系: タスクを取得できる",
			id:   1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindByID(int64(1)).Return(&entity.Task{ID: 1, Title: "タスク1", Status: entity.TaskStatusTodo, UserID: 1}, nil)
				return m
			},
			want: &entity.Task{ID: 1, Title: "タスク1", Status: entity.TaskStatusTodo, UserID: 1},
		},
		{
			name: "異常系: タスクが存在しない",
			id:   999,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindByID(int64(999)).Return(nil, errors.New("record not found"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := task.New(tt.setup(ctrl))
			got, err := uc.Get(tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{}, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("Get() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name        string
		id          int64
		title       string
		description string
		status      entity.TaskStatus
		userID      int64
		setup       func(ctrl *gomock.Controller) *repomock.MockTaskRepository
		want        *entity.Task
		wantErr     bool
	}{
		{
			name:        "正常系: タスクを更新できる",
			id:          1,
			title:       "更新タスク",
			description: "更新説明",
			status:      entity.TaskStatusInProgress,
			userID:      1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindByID(int64(1)).Return(&entity.Task{ID: 1, Title: "旧タスク", Status: entity.TaskStatusTodo, UserID: 1}, nil)
				m.EXPECT().Update(gomock.Any()).DoAndReturn(func(t *entity.Task) (*entity.Task, error) {
					return t, nil
				})
				return m
			},
			want: &entity.Task{
				ID:          1,
				Title:       "更新タスク",
				Description: "更新説明",
				Status:      entity.TaskStatusInProgress,
				UserID:      1,
			},
		},
		{
			name:        "異常系: 対象タスクが存在しない",
			id:          999,
			title:       "更新タスク",
			description: "",
			status:      entity.TaskStatusTodo,
			userID:      1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindByID(int64(999)).Return(nil, errors.New("record not found"))
				return m
			},
			wantErr: true,
		},
		{
			name:        "異常系: リポジトリの更新でエラー",
			id:          1,
			title:       "更新タスク",
			description: "",
			status:      entity.TaskStatusTodo,
			userID:      1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().FindByID(int64(1)).Return(&entity.Task{ID: 1, Title: "旧タスク", Status: entity.TaskStatusTodo, UserID: 1}, nil)
				m.EXPECT().Update(gomock.Any()).Return(nil, errors.New("db error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := task.New(tt.setup(ctrl))
			got, err := uc.Update(tt.id, tt.title, tt.description, tt.status, tt.userID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{}, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("Update() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		id      int64
		setup   func(ctrl *gomock.Controller) *repomock.MockTaskRepository
		wantErr bool
	}{
		{
			name: "正常系: タスクを削除できる",
			id:   1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().Delete(int64(1)).Return(nil)
				return m
			},
		},
		{
			name: "異常系: リポジトリエラー",
			id:   1,
			setup: func(ctrl *gomock.Controller) *repomock.MockTaskRepository {
				m := repomock.NewMockTaskRepository(ctrl)
				m.EXPECT().Delete(int64(1)).Return(errors.New("db error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := task.New(tt.setup(ctrl))
			err := uc.Delete(tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
