package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hvmidrezv/fiber-postgres/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, bookHandler *handlers.BookHandler) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 routes
	api := app.Group("/api")

	// Book routes
	books := api.Group("/books")
	books.Get("/", bookHandler.GetBooks)
	books.Post("/", bookHandler.CreateBook)
	books.Get("/:id", bookHandler.GetBookByID)
	books.Put("/:id", bookHandler.UpdateBook)
	books.Delete("/:id", bookHandler.DeleteBook)
}
