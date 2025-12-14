package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hvmidrezv/fiber-postgres/models"
	"github.com/hvmidrezv/fiber-postgres/repository"
)

type BookHandler struct {
	repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{
		repo: repo,
	}
}

// CreateBook creates a new book
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var req models.BookRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	book := &models.Book{
		Author:    req.Author,
		Title:     req.Title,
		Publisher: req.Publisher,
	}

	if err := h.repo.Create(book); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to create book",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.Response{
		Success: true,
		Message: "Book created successfully",
		Data:    book,
	})
}

// GetBooks retrieves all books
func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.repo.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{
			Error:   "Failed to fetch books",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Books fetched successfully",
		Data:    books,
	})
}

// GetBookByID retrieves a book by ID
func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request",
			Message: "Book ID is required",
		})
	}

	book, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Book not found",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Book fetched successfully",
		Data:    book,
	})
}

// UpdateBook updates an existing book
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request",
			Message: "Book ID is required",
		})
	}

	var req models.BookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
	}

	book := &models.Book{
		Author:    req.Author,
		Title:     req.Title,
		Publisher: req.Publisher,
	}

	if err := h.repo.Update(id, book); err != nil {
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Failed to update book",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Book updated successfully",
		Data:    book,
	})
}

// DeleteBook deletes a book by ID
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{
			Error:   "Invalid request",
			Message: "Book ID is required",
		})
	}

	if err := h.repo.Delete(id); err != nil {
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{
			Error:   "Failed to delete book",
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Book deleted successfully",
		Data:    nil,
	})
}
