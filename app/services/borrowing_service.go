package services

import (
	"goravel/app/models"
	"goravel/app/repositories"
)

type BorrowingService interface {
	GetAllBorrowings() ([]models.Borrowing, error)
	BorrowingUser(borrowing *models.Borrowing, userID any, bookID any) error
	ReturnUserBorrowing(borrowing *models.Borrowing, userID any, bookID any) error
	FindByUserIDBorrowing(id any) (*models.Borrowing, error)
}

type borrowingService struct {
	repo repositories.BorrowingRepository
}

func NewBorrowingService(repo repositories.BorrowingRepository) BorrowingService {
	return &borrowingService{repo: repo}
}

func (s *borrowingService) GetAllBorrowings() ([]models.Borrowing, error) {
	return s.repo.FindAllBorrowings()
}

func (s *borrowingService) BorrowingUser(borrowing *models.Borrowing, userID any, bookID any) error {
	return s.repo.BorrowingUser(borrowing, userID, bookID)
}

func (s *borrowingService) ReturnUserBorrowing(borrowing *models.Borrowing, userID any, bookID any) error {
	return s.repo.ReturnUserBorrowing(borrowing, userID, bookID)
}

func (s *borrowingService) FindByUserIDBorrowing(id any) (*models.Borrowing, error) {
	return s.repo.FindByUserIDBorrowing(id)
}
