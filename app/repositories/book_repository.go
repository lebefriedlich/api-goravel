package repositories

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type BookRepository interface {
	FindAllBook() ([]models.Book, error)
	FindByIDBook(id any) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(book *models.Book) (int64, error)
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) FindAllBook() ([]models.Book, error) {
	var books []models.Book
	err := facades.Orm().Query().Find(&books)
	return books, err
}

func (r *bookRepository) FindByIDBook(id any) (*models.Book, error) {
	var book models.Book
	err := facades.Orm().Query().Where("id", id).FirstOrFail(&book)
	return &book, err
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return facades.Orm().Query().Create(book)
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
	return facades.Orm().Query().Save(book)
}

func (r *bookRepository) DeleteBook(book *models.Book) (int64, error) {
	res, err := facades.Orm().Query().Delete(book)
	return res.RowsAffected, err
}
