package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nanch/claude-task-app/backend/domain/entity"
	"github.com/nanch/claude-task-app/backend/gen"
	userUsecase "github.com/nanch/claude-task-app/backend/usecase/user"
)

type UserHandler struct {
	usecase userUsecase.Usecase
}

func NewUserHandler(uc userUsecase.Usecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var req gen.CreateUserJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, gen.ErrorResponse{Message: err.Error()})
	}
	if req.Name == "" {
		return ctx.JSON(http.StatusBadRequest, gen.ErrorResponse{Message: "name is required"})
	}

	created, err := h.usecase.Create(req.Name, string(req.Email))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, toUserResponse(created))
}

func (h *UserHandler) UpdateUser(ctx echo.Context, userId gen.UserId) error {
	var req gen.UpdateUserJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, gen.ErrorResponse{Message: err.Error()})
	}

	updated, err := h.usecase.Update(userId, req.Name, string(req.Email))
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return ctx.JSON(http.StatusNotFound, gen.ErrorResponse{Message: "user not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, gen.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, toUserResponse(updated))
}

func toUserResponse(u *entity.User) gen.UserResponse {
	return gen.UserResponse{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
