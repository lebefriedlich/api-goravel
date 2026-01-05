package repositories

import (
	"goravel/app/models"
	"time"

	"github.com/goravel/framework/facades"
)

type BorrowingRepository interface {
	FindAllBorrowings() ([]models.Borrowing, error)
	BorrowingUser(borrowing *models.Borrowing, userID any, bookID any) error
	ReturnUserBorrowing(borrowing *models.Borrowing, userID any, bookID any) error
	FindByUserIDBorrowing(id any) (*models.Borrowing, error)
}

type borrowingRepository struct{}

func NewBorrowingRepository() BorrowingRepository {
	return &borrowingRepository{}
}

func (r *borrowingRepository) FindAllBorrowings() ([]models.Borrowing, error) {
	var borrowings []models.Borrowing
	err := facades.Orm().Query().Find(&borrowings)
	return borrowings, err
}

func (r *borrowingRepository) BorrowingUser(borrowing *models.Borrowing, userID any, bookID any) error {
	borrowing.UserID = userID.(uint)
	borrowing.BookID = bookID.(uint)
	borrowing.BorrowDate = time.Now().Format("2006-01-02 15:04:05")
	borrowing.Status = "borrowed"
	return facades.Orm().Query().Create(borrowing)
}

func (r *borrowingRepository) ReturnUserBorrowing(borrowing *models.Borrowing, userID any, bookID any) error {
	borrowing.UserID = userID.(uint)
	borrowing.BookID = bookID.(uint)

	err := facades.Orm().Query().
		Where("user_id", borrowing.UserID).
		Where("book_id", borrowing.BookID).
		First(borrowing)

	if err != nil {
		return err
	}

	borrowing.ReturnDate = time.Now().Format("2006-01-02 15:04:05")
	borrowing.Status = "returned"

	err = facades.Orm().Query().Save(borrowing)
	if err != nil {
		return err
	}

	return nil
}

func (r *borrowingRepository) FindByUserIDBorrowing(id any) (*models.Borrowing, error) {
	var borrowing models.Borrowing
	err := facades.Orm().Query().
		Where("user_id", id).
		First(&borrowing)

	if err != nil {
		return nil, err
	}
	return &borrowing, nil
}
