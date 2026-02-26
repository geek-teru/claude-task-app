package persistence_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/infrastructure/persistence"
)

func TestTaskRepository_Create(t *testing.T) {
	t.Cleanup(func() {
		truncateTable(t, "tasks")
		truncateTable(t, "users")
	})

	type args struct {
		task *entity.Task
	}

	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *entity.Task
		wantErr bool
	}{
		{
			name: "OK",
			setup: func() {
				// ユーザーだけ投入し、タスクはクリア (シーケンス競合回避)
				loadFixture(t, "testdata/tasks")
				truncateTable(t, "tasks")
			},
			args: args{task: &entity.Task{
				Title:       "新規タスク",
				Description: "新規説明",
				Status:      entity.TaskStatusTodo,
				UserID:      1,
			}},
			want: &entity.Task{
				Title:       "新規タスク",
				Description: "新規説明",
				Status:      entity.TaskStatusTodo,
				UserID:      1,
			},
			wantErr: false,
		},
	}

	repo := persistence.NewTaskRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, gotErr := repo.Create(tt.args.task)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			if got.ID == 0 {
				t.Error("expected ID to be set")
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{},
				"ID", "CreatedAt", "UpdatedAt", "DeletedAt",
			)); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTaskRepository_FindAll(t *testing.T) {
	loadFixture(t, "testdata/tasks")

	t.Cleanup(func() {
		truncateTable(t, "tasks")
		truncateTable(t, "users")
	})

	repo := persistence.NewTaskRepository(testDB)

	got, err := repo.FindAll()
	if err != nil {
		t.Fatalf("FindAll failed: %v", err)
	}

	// 論理削除済み (id=3) を除外して 2件
	if len(got) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(got))
	}

	// 内容を検証
	wantTitles := []string{"タスク1", "タスク2"}
	for i, title := range wantTitles {
		if got[i].Title != title {
			t.Errorf("tasks[%d]: expected title '%s', got '%s'", i, title, got[i].Title)
		}
	}
}

func TestTaskRepository_FindByID(t *testing.T) {
	loadFixture(t, "testdata/tasks")

	t.Cleanup(func() {
		truncateTable(t, "tasks")
		truncateTable(t, "users")
	})

	type args struct {
		id int64
	}

	tests := []struct {
		name    string
		args    args
		want    *entity.Task
		wantErr bool
	}{
		{
			name: "OK",
			args: args{id: 1},
			want: &entity.Task{
				ID:     1,
				Title:  "タスク1",
				Status: entity.TaskStatusTodo,
				UserID: 1,
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

	repo := persistence.NewTaskRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := repo.FindByID(tt.args.id)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{},
				"Description", "CreatedAt", "UpdatedAt", "DeletedAt",
			)); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTaskRepository_Update(t *testing.T) {
	t.Cleanup(func() {
		truncateTable(t, "tasks")
		truncateTable(t, "users")
	})

	type args struct {
		task *entity.Task
	}

	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *entity.Task
		wantErr bool
	}{
		{
			name: "OK",
			setup: func() {
				loadFixture(t, "testdata/tasks")
			},
			args: args{task: &entity.Task{
				ID:     1,
				Title:  "更新後タスク",
				Status: entity.TaskStatusDone,
				UserID: 1,
			}},
			want: &entity.Task{
				ID:     1,
				Title:  "更新後タスク",
				Status: entity.TaskStatusDone,
				UserID: 1,
			},
			wantErr: false,
		},
	}

	repo := persistence.NewTaskRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, gotErr := repo.Update(tt.args.task)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(entity.Task{},
				"Description", "CreatedAt", "UpdatedAt", "DeletedAt",
			)); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}

			// DB から再取得して確認
			found, _ := repo.FindByID(got.ID)
			if found.Title != tt.want.Title {
				t.Errorf("expected persisted title '%s', got '%s'", tt.want.Title, found.Title)
			}
		})
	}
}

func TestTaskRepository_Delete(t *testing.T) {
	t.Cleanup(func() {
		truncateTable(t, "tasks")
		truncateTable(t, "users")
	})

	type args struct {
		id int64
	}

	tests := []struct {
		name    string
		setup   func()
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			setup: func() {
				loadFixture(t, "testdata/tasks")
			},
			args:    args{id: 1},
			wantErr: false,
		},
	}

	repo := persistence.NewTaskRepository(testDB)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			gotErr := repo.Delete(tt.args.id)

			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("unexpected error status: gotErr=%v, wantErr=%v", gotErr, tt.wantErr)
			}
			if gotErr != nil {
				return
			}

			// 論理削除後は FindByID で見つからないことを確認
			_, err := repo.FindByID(tt.args.id)
			if err == nil {
				t.Error("expected error for deleted task")
			}
		})
	}
}
