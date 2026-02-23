package persistence_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/infrastructure/persistence"
)

func TestUserRepository_Create(t *testing.T) {
	truncateTable(t, "users")

	t.Cleanup(func() {
		truncateTable(t, "users")
	})

	type args struct {
		user *entity.User
	}

	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "OK",
			args: args{user: &entity.User{
				Name:  "新規ユーザー",
				Email: "new@example.com",
			}},
			want: &entity.User{
				Name:  "新規ユーザー",
				Email: "new@example.com",
			},
			wantErr: false,
		},
		{
			name: "ERROR: email 重複",
			setup: func() {
				loadFixture(t, "testdata/users")
			},
			args: args{user: &entity.User{
				Name:  "重複ユーザー",
				Email: "test@example.com",
			}},
			want:    nil,
			wantErr: true,
		},
	}

	repo := persistence.NewUserRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			got, gotErr := repo.Create(tt.args.user)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			if got.ID == 0 {
				t.Error("expected ID to be set")
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.User{},
				"ID", "CreatedAt", "UpdatedAt", "DeletedAt",
			)); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	loadFixture(t, "testdata/users")

	t.Cleanup(func() {
		truncateTable(t, "users")
	})

	type args struct {
		id int64
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "OK",
			args: args{id: 1},
			want: &entity.User{
				ID:    1,
				Name:  "テストユーザー",
				Email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name:    "ERROR: 存在しないID",
			args:    args{id: 99999},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "ERROR: 論理削除済み",
			args:    args{id: 3},
			want:    nil,
			wantErr: true,
		},
	}

	repo := persistence.NewUserRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := repo.FindByID(tt.args.id)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.User{},
				"CreatedAt", "UpdatedAt", "DeletedAt",
			)); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	t.Cleanup(func() {
		truncateTable(t, "users")
	})

	type args struct {
		user *entity.User
	}

	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *entity.User
		wantErr bool
	}{
		{
			name: "OK",
			setup: func() {
				loadFixture(t, "testdata/users")
			},
			args: args{user: &entity.User{
				ID:    1,
				Name:  "更新後ユーザー",
				Email: "test@example.com",
			}},
			want: &entity.User{
				ID:    1,
				Name:  "更新後ユーザー",
				Email: "test@example.com",
			},
			wantErr: false,
		},
	}

	repo := persistence.NewUserRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, gotErr := repo.Update(tt.args.user)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.User{},
				"CreatedAt", "UpdatedAt", "DeletedAt",
			)); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}

			// DB から再取得して確認
			found, _ := repo.FindByID(got.ID)
			if found.Name != tt.want.Name {
				t.Errorf("expected persisted name '%s', got '%s'", tt.want.Name, found.Name)
			}
		})
	}
}
