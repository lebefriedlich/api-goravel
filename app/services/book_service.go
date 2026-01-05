package services

import (
	"goravel/app/models"
	"goravel/app/repositories"
)

type BookService interface {
	GetAllBook() ([]models.Book, error)
	GetByIDBook(id any) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(book *models.Book) (int64, error)
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBook() ([]models.Book, error) {
	return s.repo.FindAllBook()
}

func (s *bookService) GetByIDBook(id any) (*models.Book, error) {
	return s.repo.FindByIDBook(id)
}

func (s *bookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *bookService) UpdateBook(book *models.Book) error {
	return s.repo.UpdateBook(book)
}

func (s *bookService) DeleteBook(book *models.Book) (int64, error) {
	return s.repo.DeleteBook(book)
}
