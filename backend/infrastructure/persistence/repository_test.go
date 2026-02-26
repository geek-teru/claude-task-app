package persistence_test

import (
	"fmt"
	"os"
	"testing"

	testfixtures "github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	err := setUp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()
	os.Exit(code)
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to test database: %w", err)
	}

	testDB = db
	return nil
}

func loadFixture(t *testing.T, path string) {
	t.Helper()

	sqlDB, err := testDB.DB()
	if err != nil {
		t.Fatal(err)
	}

	fixtures, err := testfixtures.New(
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(path),
		testfixtures.ResetSequencesTo(1),
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}

func truncateTable(t *testing.T, tableName string) {
	t.Helper()

	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName)
	if err := testDB.Exec(query).Error; err != nil {
		t.Fatalf("failed to truncate table %s: %v", tableName, err)
	}
}

func getTestEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
