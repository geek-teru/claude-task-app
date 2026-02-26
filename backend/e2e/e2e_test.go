package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	testfixtures "github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nanch/claude-task-app/backend/adapter/handler"
	"github.com/nanch/claude-task-app/backend/gen"
	"github.com/nanch/claude-task-app/backend/infrastructure/config"
	"github.com/nanch/claude-task-app/backend/infrastructure/persistence"
	"github.com/nanch/claude-task-app/backend/infrastructure/router"
	taskUsecase "github.com/nanch/claude-task-app/backend/usecase/task"
	userUsecase "github.com/nanch/claude-task-app/backend/usecase/user"
)

var (
	testServer *httptest.Server
	testDB     *gorm.DB
)

func TestMain(m *testing.M) {
	if err := setUp(); err != nil {
		fmt.Fprintln(os.Stderr, "E2E setup failed:", err)
		os.Exit(1)
	}
	defer testServer.Close()
	os.Exit(m.Run())
}

func setUp() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getTestEnv("DB_HOST", "localhost"),
		getTestEnv("DB_PORT", "5432"),
		getTestEnv("DB_USER", "app"),
		getTestEnv("DB_PASSWORD", "password"),
		getTestEnv("DB_NAME", "claude_task_app"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	testDB = db

	// マイグレーション実行
	if err := config.Migrate(db); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	// DI 組み立て
	taskRepo := persistence.NewTaskRepository(db)
	userRepo := persistence.NewUserRepository(db)
	taskUC := taskUsecase.New(taskRepo)
	userUC := userUsecase.New(userRepo)
	taskHandler := handler.NewTaskHandler(taskUC)
	userHandler := handler.NewUserHandler(userUC)
	h := handler.NewHandler(taskHandler, userHandler)
	e := router.New(h)

	testServer = httptest.NewServer(e)
	return nil
}

func loadFixture(t *testing.T) {
	t.Helper()
	sqlDB, err := testDB.DB()
	if err != nil {
		t.Fatal(err)
	}
	fixtures, err := testfixtures.New(
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("testdata"),
		testfixtures.ResetSequencesTo(10),
	)
	if err != nil {
		t.Fatal(err)
	}
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}

func getTestEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// ===== ユーザー =====

func TestE2E_CreateUser(t *testing.T) {
	loadFixture(t)
	t.Cleanup(func() { truncate(t, "users", "tasks") })

	tests := []struct {
		name       string
		body       map[string]any
		wantStatus int
	}{
		{
			name:       "正常系: ユーザーを作成できる",
			body:       map[string]any{"name": "新規ユーザー", "email": "new@example.com"},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "異常系: name が空",
			body:       map[string]any{"email": "no-name@example.com"},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.body)
			resp, err := http.Post(testServer.URL+"/api/v1/users", "application/json", bytes.NewReader(b))
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("CreateUser status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestE2E_UpdateUser(t *testing.T) {
	loadFixture(t)
	t.Cleanup(func() { truncate(t, "users", "tasks") })

	tests := []struct {
		name       string
		userID     int
		body       map[string]any
		wantStatus int
	}{
		{
			name:       "正常系: ユーザーを更新できる",
			userID:     1,
			body:       map[string]any{"name": "更新ユーザー", "email": "updated@example.com"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "異常系: 存在しないユーザー",
			userID:     999,
			body:       map[string]any{"name": "更新ユーザー", "email": "updated@example.com"},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/users/%d", testServer.URL, tt.userID), bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("UpdateUser status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

// ===== タスク =====

func TestE2E_CreateTask(t *testing.T) {
	loadFixture(t)
	t.Cleanup(func() { truncate(t, "users", "tasks") })

	tests := []struct {
		name       string
		body       map[string]any
		wantStatus int
	}{
		{
			name:       "正常系: タスクを作成できる",
			body:       map[string]any{"title": "新規タスク", "userId": 1},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "異常系: title が空",
			body:       map[string]any{"userId": 1},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.body)
			resp, err := http.Post(testServer.URL+"/api/v1/tasks", "application/json", bytes.NewReader(b))
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("CreateTask status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestE2E_ListTasks(t *testing.T) {
	loadFixture(t)

	resp, err := http.Get(testServer.URL + "/api/v1/tasks")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("ListTasks status = %d, want 200", resp.StatusCode)
	}

	var tasks []gen.TaskResponse
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Errorf("ListTasks count = %d, want 2", len(tasks))
	}
}

func TestE2E_GetTask(t *testing.T) {
	loadFixture(t)

	tests := []struct {
		name       string
		taskID     int
		wantStatus int
	}{
		{
			name:       "正常系: タスクを取得できる",
			taskID:     1,
			wantStatus: http.StatusOK,
		},
		{
			name:       "異常系: 存在しないタスク",
			taskID:     999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/api/v1/tasks/%d", testServer.URL, tt.taskID))
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("GetTask status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestE2E_UpdateTask(t *testing.T) {
	loadFixture(t)
	t.Cleanup(func() { truncate(t, "users", "tasks") })

	tests := []struct {
		name       string
		taskID     int
		body       map[string]any
		wantStatus int
	}{
		{
			name:       "正常系: タスクを更新できる",
			taskID:     1,
			body:       map[string]any{"title": "更新タスク", "status": "done", "userId": 1},
			wantStatus: http.StatusOK,
		},
		{
			name:       "異常系: 存在しないタスク",
			taskID:     999,
			body:       map[string]any{"title": "更新タスク", "userId": 1},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/v1/tasks/%d", testServer.URL, tt.taskID), bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("UpdateTask status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestE2E_DeleteTask(t *testing.T) {
	loadFixture(t)
	t.Cleanup(func() { truncate(t, "users", "tasks") })

	tests := []struct {
		name       string
		taskID     int
		wantStatus int
	}{
		{
			name:       "正常系: タスクを削除できる",
			taskID:     1,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "異常系: 存在しないタスク (論理削除済み or 未存在)",
			taskID:     999,
			wantStatus: http.StatusNoContent, // Delete は存在確認しない実装のため 204
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/tasks/%d", testServer.URL, tt.taskID), nil)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("DeleteTask status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func truncate(t *testing.T, tables ...string) {
	t.Helper()
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)
		if err := testDB.Exec(query).Error; err != nil {
			t.Logf("failed to truncate %s: %v", table, err)
		}
	}
}
