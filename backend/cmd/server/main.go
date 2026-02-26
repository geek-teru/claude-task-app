package main

import (
	"log"

	"github.com/nanch/claude-task-app/backend/adapter/handler"
	"github.com/nanch/claude-task-app/backend/infrastructure/config"
	"github.com/nanch/claude-task-app/backend/infrastructure/persistence"
	"github.com/nanch/claude-task-app/backend/infrastructure/router"
	taskUsecase "github.com/nanch/claude-task-app/backend/usecase/task"
	userUsecase "github.com/nanch/claude-task-app/backend/usecase/user"
)

func main() {
	// DB 接続
	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// リポジトリ
	taskRepo := persistence.NewTaskRepository(db)
	userRepo := persistence.NewUserRepository(db)

	// ユースケース
	taskUC := taskUsecase.New(taskRepo)
	userUC := userUsecase.New(userRepo)

	// ハンドラー
	taskHandler := handler.NewTaskHandler(taskUC)
	userHandler := handler.NewUserHandler(userUC)
	h := handler.NewHandler(taskHandler, userHandler)

	// ルーター起動
	e := router.New(h)
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
