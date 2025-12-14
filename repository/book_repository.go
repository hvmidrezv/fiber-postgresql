package repository

import (
	"github.com/hvmidrezv/fiber-postgres/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id string) (*models.Book, error)
	Update(id string, book *models.Book) error
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id string) (*models.Book, error) {
	var book models.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Update(id string, book *models.Book) error {
	var existingBook models.Book
	if err := r.db.Where("id = ?", id).First(&existingBook).Error; err != nil {
		return err
	}

	existingBook.Author = book.Author
	existingBook.Title = book.Title
	existingBook.Publisher = book.Publisher

	return r.db.Save(&existingBook).Error
}

func (r *bookRepository) Delete(id string) error {
	var book models.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}
	return r.db.Delete(&book).Error
}
