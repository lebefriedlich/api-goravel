package helpers

import (
	"goravel/app/models"
)

type BorrowingResponse map[string]any

func ToBorrowingResponse(borrowing *models.Borrowing) BorrowingResponse {
	return BorrowingResponse{
		"id":      borrowing.ID,
		"user_id": borrowing.UserID,
		"book_id": borrowing.BookID,
		"status":  borrowing.Status,
	}
}

func ToBorrowingResponseList(borrowings []models.Borrowing) []BorrowingResponse {
	var response []BorrowingResponse
	for _, borrowing := range borrowings {
		response = append(response, ToBorrowingResponse(&borrowing))
	}
	return response
}
