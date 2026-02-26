package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/gen"
	taskUsecase "github.com/nanch/claude-task-app/backend/usecase/task"
)

type TaskHandler struct {
	usecase taskUsecase.Usecase
}

func NewTaskHandler(uc taskUsecase.Usecase) *TaskHandler {
	return &TaskHandler{usecase: uc}
}

func (h *TaskHandler) CreateTask(ctx echo.Context) error {
	var req gen.CreateTaskJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, gen.ErrorResponse{Message: err.Error()})
	}
	if req.Title == "" {
		return ctx.JSON(http.StatusBadRequest, gen.ErrorResponse{Message: "title is required"})
	}

	description := ""
	if req.Description != nil {
		description = *req.Description
	}
	status := entity.TaskStatusTodo
	if req.Status != nil {
		status = entity.TaskStatus(*req.Status)
	}

	created, err := h.usecase.Create(req.Title, description, status, req.UserId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, toTaskResponse(created))
}

func (h *TaskHandler) ListTasks(ctx echo.Context) error {
	tasks, err := h.usecase.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}

	resp := make([]gen.TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		resp = append(resp, toTaskResponse(t))
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (h *TaskHandler) GetTask(ctx echo.Context, taskId gen.TaskId) error {
	t, err := h.usecase.Get(taskId)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, gen.ErrorResponse{Message: "task not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, toTaskResponse(t))
}

func (h *TaskHandler) UpdateTask(ctx echo.Context, taskId gen.TaskId) error {
	var req gen.UpdateTaskJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, gen.ErrorResponse{Message: err.Error()})
	}

	description := ""
	if req.Description != nil {
		description = *req.Description
	}
	status := entity.TaskStatusTodo
	if req.Status != nil {
		status = entity.TaskStatus(*req.Status)
	}

	updated, err := h.usecase.Update(taskId, req.Title, description, status, req.UserId)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, gen.ErrorResponse{Message: "task not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, toTaskResponse(updated))
}

func (h *TaskHandler) DeleteTask(ctx echo.Context, taskId gen.TaskId) error {
	// Delete は冪等操作: 存在確認せず削除し常に 204 を返す
	if err := h.usecase.Delete(taskId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

func toTaskResponse(t *entity.Task) gen.TaskResponse {
	desc := t.Description
	return gen.TaskResponse{
		Id:          t.ID,
		Title:       t.Title,
		Description: &desc,
		Status:      gen.TaskStatus(t.Status),
		UserId:      t.UserID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
