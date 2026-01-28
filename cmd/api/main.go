package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go-api-boilerplate/handler"
	"go-api-boilerplate/internal/config"
	"go-api-boilerplate/internal/db"
	"go-api-boilerplate/internal/server"
	"go-api-boilerplate/pkg/common/logger"
)

func main() {
	ctx := context.Background()
	commonInit()

	// Initialize the server instance
	svr := server.NewServer()
	if err := svr.RegisterService(svr.GetEcho(), handler.Handler{}); err != nil {
		logger.WithContext(ctx).Error("Failed to register service", "error", err)
		panic(err)
	}
	
	// Run graceful shutdown in a separate goroutine
	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)
	go gracefulShutdown(svr, done)

	// Run the server
	err := svr.Run()
	if err != nil && err != http.ErrServerClosed {
		logger.WithContext(ctx).Error("Failed to start server", "error", err)
	}

	// Wait for the graceful shutdown to complete
	<-done
	logger.WithContext(ctx).Info("Graceful shutdown complete.")
}

// Add inits here
func commonInit() {
	logger.Init()
	config.InitEnv()
	db.InitDB()
}

// gracefulShutdown handles graceful shutdown of the server when the user presses Ctrl+C via os signal context
func gracefulShutdown(s server.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	logger.Logger.Info("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := s.Stop(ctx); err != nil {
		logger.Logger.Error("Server forced to shutdown with error", "error", err)
	}

	// // Close the database connection
	// s.DB.Close()

	logger.Logger.Info("Server exiting")
	// Notify the main goroutine that the shutdown is complete
	done <- true
}