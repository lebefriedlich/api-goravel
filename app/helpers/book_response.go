package helpers

import (
	"goravel/app/models"
)

type BookResponse map[string]any

func ToBookResponse(book *models.Book) BookResponse {
	return BookResponse{
		"id":             book.ID,
		"author":         book.Author,
		"title":          book.Title,
		"published_year": book.PublishedYear,
		"stock":          book.Stock,
	}
}

func ToBookResponseList(books []models.Book) []BookResponse {
	var response []BookResponse
	for _, book := range books {
		response = append(response, ToBookResponse(&book))
	}
	return response
}
