package models

import (
	"time"

	"gorm.io/gorm"
)

// Book represents a book entity in the database
type Book struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Author    string         `gorm:"type:varchar(255);not null" json:"author"`
	Title     string         `gorm:"type:varchar(255);not null;index" json:"title"`
	Publisher string         `gorm:"type:varchar(255)" json:"publisher"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BookRequest represents the request body for creating/updating a book
type BookRequest struct {
	Author    string `json:"author" validate:"required,min=1,max=255"`
	Title     string `json:"title" validate:"required,min=1,max=255"`
	Publisher string `json:"publisher" validate:"max=255"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// MigrateBooks runs database migrations for the Book model
func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}

// TableName specifies the table name for the Book model
func (Book) TableName() string {
	return "books"
}
