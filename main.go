package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/hvmidrezv/fiber-postgres/config"
	"github.com/hvmidrezv/fiber-postgres/handlers"
	"github.com/hvmidrezv/fiber-postgres/middleware"
	"github.com/hvmidrezv/fiber-postgres/models"
	"github.com/hvmidrezv/fiber-postgres/repository"
	"github.com/hvmidrezv/fiber-postgres/routes"
	"github.com/hvmidrezv/fiber-postgres/storage"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := storage.NewConnection(cfg.Database.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := storage.Close(db); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Run migrations
	if err := models.MigrateBooks(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully")

	// Initialize repository
	bookRepo := repository.NewBookRepository(db)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookRepo)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Fiber-PostgreSQL API",
		ErrorHandler: customErrorHandler,
	})

	// Setup middleware
	middleware.SetupMiddleware(app)

	// Setup routes
	routes.SetupRoutes(app, bookHandler)

	// Start server in a goroutine
	go func() {
		serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
		log.Printf("Server starting on http://localhost%s", serverAddr)
		if err := app.Listen(serverAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
	log.Println("Server stopped gracefully")
}

// customErrorHandler handles errors globally
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(models.ErrorResponse{
		Error:   message,
		Message: err.Error(),
	})
}
