package server

import (
	"go-api-boilerplate/handler"

	"github.com/labstack/echo/v4"
)

func RegisterGrootRoutes(e *echo.Echo, h *handler.Handler) {

	e.GET("/", handler.HandleGroot(h))
}

func RegisterUserRoutes(e *echo.Echo, h *handler.Handler) {

	e.POST("/users", handler.HandleCreateUser(h))
	e.GET("/users", handler.HandleListUsers(h))
}
