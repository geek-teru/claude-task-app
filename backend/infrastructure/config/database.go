package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nanch/claude-task-app/backend/infrastructure/persistence"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "app"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "claude_task_app"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// 既存の外部キー制約を削除
	if err := dropAllForeignKeys(db); err != nil {
		return err
	}

	return db.AutoMigrate(
		&persistence.UserModel{},
		&persistence.TaskModel{},
	)
}

func dropAllForeignKeys(db *gorm.DB) error {
	type fk struct {
		ConstraintName string
		TableName      string
	}

	var fks []fk
	err := db.Raw(`
		SELECT tc.constraint_name, tc.table_name
		FROM information_schema.table_constraints tc
		WHERE tc.constraint_type = 'FOREIGN KEY'
		  AND tc.table_schema = 'public'
	`).Scan(&fks).Error
	if err != nil {
		return fmt.Errorf("failed to query foreign keys: %w", err)
	}

	for _, f := range fks {
		query := fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", f.TableName, f.ConstraintName)
		if err := db.Exec(query).Error; err != nil {
			return fmt.Errorf("failed to drop FK %s: %w", f.ConstraintName, err)
		}
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
