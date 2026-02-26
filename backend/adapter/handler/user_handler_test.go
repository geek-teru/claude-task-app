package handler_test

import (
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
	usermock "github.com/nanch/claude-task-app/backend/usecase/user/mock"
)

func TestCreateUser(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		body       string
		setup      func() *usermock.UserUsecase
		wantStatus int
	}{
		{
			name: "正常系: ユーザーを作成できる",
			body: `{"name":"テストユーザー","email":"test@example.com"}`,
			setup: func() *usermock.UserUsecase {
				return &usermock.UserUsecase{
					CreateFn: func(name, email string) (*entity.User, error) {
						return &entity.User{ID: 1, Name: name, Email: email, CreatedAt: now, UpdatedAt: now}, nil
					},
				}
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "異常系: 不正なJSON",
			body:       `{invalid}`,
			setup:      func() *usermock.UserUsecase { return &usermock.UserUsecase{} },
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "異常系: usecase エラー",
			body: `{"name":"テストユーザー","email":"test@example.com"}`,
			setup: func() *usermock.UserUsecase {
				return &usermock.UserUsecase{
					CreateFn: func(name, email string) (*entity.User, error) {
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
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(tt.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := handler.NewUserHandler(tt.setup())
			_ = h.CreateUser(ctx)

			if rec.Code != tt.wantStatus {
				t.Errorf("CreateUser() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		userID     int64
		body       string
		setup      func() *usermock.UserUsecase
		wantStatus int
	}{
		{
			name:   "正常系: ユーザーを更新できる",
			userID: 1,
			body:   `{"name":"更新ユーザー","email":"updated@example.com"}`,
			setup: func() *usermock.UserUsecase {
				return &usermock.UserUsecase{
					UpdateFn: func(id int64, name, email string) (*entity.User, error) {
						return &entity.User{ID: id, Name: name, Email: email, CreatedAt: now, UpdatedAt: now}, nil
					},
				}
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "異常系: ユーザーが存在しない",
			userID: 999,
			body:   `{"name":"更新ユーザー","email":"updated@example.com"}`,
			setup: func() *usermock.UserUsecase {
				return &usermock.UserUsecase{
					UpdateFn: func(id int64, name, email string) (*entity.User, error) {
						return nil, fmt.Errorf("not found: %w", entity.ErrNotFound)
					},
				}
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "異常系: 不正なJSON",
			userID:     1,
			body:       `{invalid}`,
			setup:      func() *usermock.UserUsecase { return &usermock.UserUsecase{} },
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

			h := handler.NewUserHandler(tt.setup())
			_ = h.UpdateUser(ctx, tt.userID)

			if rec.Code != tt.wantStatus {
				t.Errorf("UpdateUser() status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
