package handler

import "github.com/nanch/claude-task-app/backend/gen"

// Handler は gen.ServerInterface を実装する統合ハンドラー
type Handler struct {
	*TaskHandler
	*UserHandler
}

// NewHandler は全ハンドラーを束ねた Handler を返す
func NewHandler(task *TaskHandler, user *UserHandler) gen.ServerInterface {
	return &Handler{
		TaskHandler: task,
		UserHandler: user,
	}
}
