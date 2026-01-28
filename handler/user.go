package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-api-boilerplate/pkg/common/logger"
	"go-api-boilerplate/pkg/domain/model"
	httpsuite "go-api-boilerplate/pkg/http"
	"go-api-boilerplate/service"
)

func HandleCreateUser(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		req, err := httpsuite.ParseRequest[model.CreateUserRequest](w, c.Request())
		if err != nil {
			logger.Logger.Error("Failed to parse request", "error", err)
			return err
		}
		logger.Logger.Info("Received a request to create a user", "request", req)

		err = service.CreateUser(req.Name, req.Role)
		if err != nil {
			return err
		}

		data := map[string]any{
			"message": "success",
		}
		httpsuite.SendResponse(w, http.StatusOK, data)
		return nil
	}
}

func HandleListUsers(h *Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		logger.Logger.Info("Received a request to list all users", "request", c.Request())

		users, err := service.ListUsers()
		if err != nil {
			return err
		}

		data := map[string]any{
			"message": "success",
			"users":   users,
		}

		httpsuite.SendResponse(w, http.StatusOK, data)
		return nil
	}
}	
