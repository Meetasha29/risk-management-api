package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"risky_plumber/handler"
	"risky_plumber/repository"
	"risky_plumber/router"
	"risky_plumber/service"
)

func main() {
	log.Println(" Risky Plumber API is starting on port 8080...")

	// Initialize dependencies
	riskRepo := repository.NewInMemoryRiskRepo()
	riskEngine := service.NewEngine(riskRepo)
	riskHandler := handler.NewRiskHandler(riskEngine)

	// Setup routes using router
	mux := router.SetupRoutes(riskHandler)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s", err)
		}
	}()

	log.Println("Server is running on http://localhost:8080")

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %s", err)
	}

	log.Println("✅ Server gracefully stopped")
}
