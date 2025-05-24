package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iniudin/boilerplate/internal/platform"
	"github.com/iniudin/boilerplate/internal/server"
)

// @title Boilerplate API
// @version 1.0
// @description Boilerplate API
// @host localhost:8080
// @BasePath /
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := platform.NewConfig()

	pool, err := platform.NewPostgres(ctx, config)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Database ping failed:", err)
		return
	}
	fmt.Println("Database connected")

	server := &http.Server{
		Addr:    config.APP_HOST + ":" + config.APP_PORT,
		Handler: server.NewRouter(config),
	}

	go func() {
		fmt.Println("Server is running on http://" + config.APP_HOST + ":" + config.APP_PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("HTTP server error:", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Server shutdown error:", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}
}
