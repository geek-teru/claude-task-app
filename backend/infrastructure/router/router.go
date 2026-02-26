package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nanch/claude-task-app/backend/gen"
)

func New(h gen.ServerInterface) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	gen.RegisterHandlers(e, h)

	return e
}
