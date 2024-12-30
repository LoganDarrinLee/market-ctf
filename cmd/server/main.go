package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LoganDarrinLee/market-ctf/internal/common"
	"github.com/LoganDarrinLee/market-ctf/internal/config"
	"github.com/LoganDarrinLee/market-ctf/internal/middleware"
	"github.com/LoganDarrinLee/market-ctf/internal/routing"
)

func main() {
	// New environment configuration
	env := config.NewEnv()

	// New context
	ctx := context.Background()

	// Basic Logging
	logger := &common.BasicLogger{}

	// Base routing handler
	appHandler := routing.NewHandler(logger)

	// Setup the base handlers database pool
	appHandler.Pool = config.InitDB(ctx, env)
	defer appHandler.Pool.Close()

	// Router setup
	router := http.NewServeMux()

	// Configure base routes
	routing.ConfigureRoutes(router, appHandler)

	// Create middleware stack
	stack := middleware.CreateStack(
		middleware.WithRequestContext,
		middleware.CheckAuth,
		middleware.Logging,
	)

	// Setup server
	server := http.Server{
		Addr:    env.ServerPort,
		Handler: stack(router),
	}

	// Channel for graceful shutdowns, will finish handling current requests.
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		shutDownCtx, shutDownRelease := context.WithTimeout(ctx, 10*time.Second)
		defer shutDownRelease()

		if err := server.Shutdown(shutDownCtx); err != nil {
			log.Fatalf("HTTP shutdown error: %v", err)
		}
		log.Println("Graceful shutdown complete.")
	}()

	// Runserver
	log.Println("Starting sever on port: ", env.ServerPort)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP Server error: %v", err)
	}

	// Successful shutdown.
	log.Println("Stopped serving HTTP connections.")
}
