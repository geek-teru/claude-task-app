package entity

import "time"

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          int64
	Title       string
	Description string
	Status      TaskStatus
	UserID      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
