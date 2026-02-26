package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nanch/claude-task-app/backend/adapter/handler"
	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/gen"
	taskmock "github.com/nanch/claude-task-app/backend/usecase/task/mock"
)

func TestCreateTask(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		body       string
		setup      func() *taskmock.TaskUsecase
		wantStatus int
		wantID     int64
	}{
		{
			name: "正常系: タスクを作成できる",
			body: `{"title":"テストタスク","description":"説明","userId":1}`,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					CreateFn: func(title, desc string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
						return &entity.Task{ID: 1, Title: title, Description: desc, Status: status, UserID: userID, CreatedAt: now, UpdatedAt: now}, nil
					},
				}
			},
			wantStatus: http.StatusCreated,
			wantID:     1,
		},
		{
			name:       "異常系: 不正なJSON",
			body:       `{invalid}`,
			setup:      func() *taskmock.TaskUsecase { return &taskmock.TaskUsecase{} },
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "異常系: usecase エラー",
			body: `{"title":"テストタスク","userId":1}`,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					CreateFn: func(title, desc string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
						return nil, errors.New("db error")
					},
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := handler.NewTaskHandler(tt.setup())
			_ = h.CreateTask(ctx)

			if rec.Code != tt.wantStatus {
				t.Errorf("CreateTask() status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantID > 0 {
				var resp gen.TaskResponse
				_ = json.Unmarshal(rec.Body.Bytes(), &resp)
				if resp.Id != tt.wantID {
					t.Errorf("CreateTask() id = %d, want %d", resp.Id, tt.wantID)
				}
			}
		})
	}
}

func TestListTasks(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		setup      func() *taskmock.TaskUsecase
		wantStatus int
		wantCount  int
	}{
		{
			name: "正常系: タスク一覧を取得できる",
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					ListFn: func() ([]*entity.Task, error) {
						return []*entity.Task{
							{ID: 1, Title: "タスク1", Status: entity.TaskStatusTodo, UserID: 1, CreatedAt: now, UpdatedAt: now},
							{ID: 2, Title: "タスク2", Status: entity.TaskStatusDone, UserID: 1, CreatedAt: now, UpdatedAt: now},
						}, nil
					},
				}
			},
			wantStatus: http.StatusOK,
			wantCount:  2,
		},
		{
			name: "正常系: タスクが0件",
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					ListFn: func() ([]*entity.Task, error) {
						return []*entity.Task{}, nil
					},
				}
			},
			wantStatus: http.StatusOK,
			wantCount:  0,
		},
		{
			name: "異常系: usecase エラー",
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					ListFn: func() ([]*entity.Task, error) {
						return nil, errors.New("db error")
					},
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := handler.NewTaskHandler(tt.setup())
			_ = h.ListTasks(ctx)

			if rec.Code != tt.wantStatus {
				t.Errorf("ListTasks() status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantCount > 0 {
				var resp []gen.TaskResponse
				_ = json.Unmarshal(rec.Body.Bytes(), &resp)
				if len(resp) != tt.wantCount {
					t.Errorf("ListTasks() count = %d, want %d", len(resp), tt.wantCount)
				}
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		taskID     int64
		setup      func() *taskmock.TaskUsecase
		wantStatus int
	}{
		{
			name:   "正常系: タスクを取得できる",
			taskID: 1,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					GetFn: func(id int64) (*entity.Task, error) {
						return &entity.Task{ID: id, Title: "タスク1", Status: entity.TaskStatusTodo, UserID: 1, CreatedAt: now, UpdatedAt: now}, nil
					},
				}
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "異常系: タスクが存在しない",
			taskID: 999,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					GetFn: func(id int64) (*entity.Task, error) {
						return nil, fmt.Errorf("not found: %w", entity.ErrNotFound)
					},
				}
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "異常系: usecase エラー",
			taskID: 1,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					GetFn: func(id int64) (*entity.Task, error) {
						return nil, errors.New("db error")
					},
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := handler.NewTaskHandler(tt.setup())
			_ = h.GetTask(ctx, tt.taskID)

			if rec.Code != tt.wantStatus {
				t.Errorf("GetTask() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		taskID     int64
		body       string
		setup      func() *taskmock.TaskUsecase
		wantStatus int
	}{
		{
			name:   "正常系: タスクを更新できる",
			taskID: 1,
			body:   `{"title":"更新タスク","userId":1}`,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					UpdateFn: func(id int64, title, desc string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
						return &entity.Task{ID: id, Title: title, Status: status, UserID: userID, CreatedAt: now, UpdatedAt: now}, nil
					},
				}
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "異常系: タスクが存在しない",
			taskID: 999,
			body:   `{"title":"更新タスク","userId":1}`,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					UpdateFn: func(id int64, title, desc string, status entity.TaskStatus, userID int64) (*entity.Task, error) {
						return nil, fmt.Errorf("not found: %w", entity.ErrNotFound)
					},
				}
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "異常系: 不正なJSON",
			taskID:     1,
			body:       `{invalid}`,
			setup:      func() *taskmock.TaskUsecase { return &taskmock.TaskUsecase{} },
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := handler.NewTaskHandler(tt.setup())
			_ = h.UpdateTask(ctx, tt.taskID)

			if rec.Code != tt.wantStatus {
				t.Errorf("UpdateTask() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name       string
		taskID     int64
		setup      func() *taskmock.TaskUsecase
		wantStatus int
	}{
		{
			name:   "正常系: タスクを削除できる",
			taskID: 1,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					DeleteFn: func(id int64) error { return nil },
				}
			},
			wantStatus: http.StatusNoContent,
		},
		{
			name:   "異常系: usecase エラー",
			taskID: 1,
			setup: func() *taskmock.TaskUsecase {
				return &taskmock.TaskUsecase{
					DeleteFn: func(id int64) error { return errors.New("db error") },
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := handler.NewTaskHandler(tt.setup())
			_ = h.DeleteTask(ctx, tt.taskID)

			if rec.Code != tt.wantStatus {
				t.Errorf("DeleteTask() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
