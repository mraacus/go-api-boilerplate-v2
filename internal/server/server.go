package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-api-boilerplate/handler"
	"go-api-boilerplate/internal/db"
	"go-api-boilerplate/pkg/common/logger"
	"go-api-boilerplate/pkg/middleware"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Run() 													error
	Stop(ctx context.Context) 								error
	RegisterService(e *echo.Echo, handler handler.Handler)  error
	GetEcho() 												*echo.Echo
}

type server struct {
	echo *echo.Echo
}

func NewServer() Server {
	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.CustomMiddleware)

	return &server{
		echo: e,
	}
}

func (s *server) Run() (err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // Default port
	}
	logger.Logger.Info("Server starting up and listening on port:", "port", port)

	// Configure server timeouts
	s.echo.Server.ReadTimeout = 10 * time.Second
	s.echo.Server.WriteTimeout = 30 * time.Second
	s.echo.Server.IdleTimeout = time.Minute

	// Start the Echo server
	return s.echo.Start(fmt.Sprintf(":%d", port))
}

func (s *server) Stop(ctx context.Context) (err error) {
	return s.echo.Shutdown(ctx)
}

func (s *server) GetEcho() *echo.Echo {
	return s.echo
}


func (s *server) RegisterService(e *echo.Echo, handler handler.Handler) error {
	if e == nil {
		panic("echo instance is nil, please specify a non nil echo instance")
	}

	// Set up custom CORS middleware for the Echo instance
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Database health check route
	e.GET("/health", func(c echo.Context) error {
		healthData := db.Health(db.GetReadDB())
		return c.JSON(http.StatusOK, healthData)
	})

	// Register routes
	RegisterGrootRoutes(e, &handler)
	RegisterUserRoutes(e, &handler)

	return nil
}

