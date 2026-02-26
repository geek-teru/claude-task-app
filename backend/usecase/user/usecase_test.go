package user_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nanch/claude-task-app/backend/domain/entity"
	repomock "github.com/nanch/claude-task-app/backend/domain/repository/mock"
	"github.com/nanch/claude-task-app/backend/usecase/user"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name    string
		uname   string
		email   string
		setup   func(ctrl *gomock.Controller) *repomock.MockUserRepository
		want    *entity.User
		wantErr bool
	}{
		{
			name:  "正常系: ユーザーを作成できる",
			uname: "テストユーザー",
			email: "test@example.com",
			setup: func(ctrl *gomock.Controller) *repomock.MockUserRepository {
				m := repomock.NewMockUserRepository(ctrl)
				m.EXPECT().Create(gomock.Any()).DoAndReturn(func(u *entity.User) (*entity.User, error) {
					u.ID = 1
					return u, nil
				})
				return m
			},
			want: &entity.User{
				ID:    1,
				Name:  "テストユーザー",
				Email: "test@example.com",
			},
		},
		{
			name:  "異常系: リポジトリエラー",
			uname: "テストユーザー",
			email: "test@example.com",
			setup: func(ctrl *gomock.Controller) *repomock.MockUserRepository {
				m := repomock.NewMockUserRepository(ctrl)
				m.EXPECT().Create(gomock.Any()).Return(nil, errors.New("db error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := user.New(tt.setup(ctrl))
			got, err := uc.Create(tt.uname, tt.email)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("Create() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name    string
		id      int64
		uname   string
		email   string
		setup   func(ctrl *gomock.Controller) *repomock.MockUserRepository
		want    *entity.User
		wantErr bool
	}{
		{
			name:  "正常系: ユーザーを更新できる",
			id:    1,
			uname: "更新ユーザー",
			email: "updated@example.com",
			setup: func(ctrl *gomock.Controller) *repomock.MockUserRepository {
				m := repomock.NewMockUserRepository(ctrl)
				m.EXPECT().FindByID(int64(1)).Return(&entity.User{ID: 1, Name: "旧ユーザー", Email: "old@example.com"}, nil)
				m.EXPECT().Update(gomock.Any()).DoAndReturn(func(u *entity.User) (*entity.User, error) {
					return u, nil
				})
				return m
			},
			want: &entity.User{
				ID:    1,
				Name:  "更新ユーザー",
				Email: "updated@example.com",
			},
		},
		{
			name:  "異常系: 対象ユーザーが存在しない",
			id:    999,
			uname: "更新ユーザー",
			email: "updated@example.com",
			setup: func(ctrl *gomock.Controller) *repomock.MockUserRepository {
				m := repomock.NewMockUserRepository(ctrl)
				m.EXPECT().FindByID(int64(999)).Return(nil, errors.New("record not found"))
				return m
			},
			wantErr: true,
		},
		{
			name:  "異常系: リポジトリの更新でエラー",
			id:    1,
			uname: "更新ユーザー",
			email: "updated@example.com",
			setup: func(ctrl *gomock.Controller) *repomock.MockUserRepository {
				m := repomock.NewMockUserRepository(ctrl)
				m.EXPECT().FindByID(int64(1)).Return(&entity.User{ID: 1, Name: "旧ユーザー", Email: "old@example.com"}, nil)
				m.EXPECT().Update(gomock.Any()).Return(nil, errors.New("db error"))
				return m
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			uc := user.New(tt.setup(ctrl))
			got, err := uc.Update(tt.id, tt.uname, tt.email)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
				t.Errorf("Update() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
