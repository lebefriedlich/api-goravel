package controllers

import (
	"goravel/app/helpers"
	"goravel/app/models"
	"goravel/app/repositories"
	"goravel/app/services"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

type BookController struct {
	service services.BookService
}

func NewBookController() *BookController {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)
	return &BookController{service: service}
}

func (r *BookController) Index(ctx http.Context) http.Response {
	books, err := r.service.GetAllBook()
	if err != nil {
		return helpers.Error(ctx, 500, "Failed to fetch books", err.Error())
	}

	bookResponses := helpers.ToBookResponseList(books)

	return helpers.Success(ctx, "Books retrieved successfully", bookResponses)
}

func (r *BookController) Show(ctx http.Context) http.Response {
	book, err := r.service.GetByIDBook(ctx.Request().Input("id"))
	if err != nil {
		return helpers.Error(ctx, 404, "Book not found", err.Error())
	}

	return helpers.Success(ctx, "Book retrieved successfully", helpers.ToBookResponse(book))
}

func (r *BookController) Store(ctx http.Context) http.Response {
	validation, err := ctx.Request().Validate(map[string]string{
		"author":         "required|string|max_len:255",
		"title":          "required|string|max_len:255",
		"published_year": "required|integer",
		"stock":          "required|integer",
	})

	if err != nil {
		return helpers.Error(ctx, 500, "Validation setup failed", err.Error())
	}

	if validation.Fails() {
		return helpers.Error(ctx, 400, "Validation failed", validation.Errors().All())
	}

	stockStr := ctx.Request().Input("stock")
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return helpers.Error(ctx, 400, "Invalid stock value", err.Error())
	}

	publishedYearStr := ctx.Request().Input("published_year")
	publishedYear, err := strconv.Atoi(publishedYearStr)
	if err != nil {
		return helpers.Error(ctx, 400, "Invalid published year value", err.Error())
	}

	book := &models.Book{
		Author:        ctx.Request().Input("author"),
		Title:         ctx.Request().Input("title"),
		PublishedYear: publishedYear,
		Stock:         stock,
	}

	if err := r.service.CreateBook(book); err != nil {
		return helpers.Error(ctx, 500, "Failed to create book", err.Error())
	}

	return helpers.Success(ctx, "Book created successfully", helpers.ToBookResponse(book))
}

func (r *BookController) Update(ctx http.Context) http.Response {
	validation, err := ctx.Request().Validate(map[string]string{
		"author":         "required|string|max_len:255",
		"title":          "required|string|max_len:255",
		"published_year": "required|integer",
		"stock":          "required|integer",
	})

	if err != nil {
		return helpers.Error(ctx, 500, "Validation setup failed", err.Error())
	}

	if validation.Fails() {
		return helpers.Error(ctx, 400, "Validation failed", validation.Errors().All())
	}

	stockStr := ctx.Request().Input("stock")
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return helpers.Error(ctx, 400, "Invalid stock value", err.Error())
	}

	publishedYearStr := ctx.Request().Input("published_year")
	publishedYear, err := strconv.Atoi(publishedYearStr)
	if err != nil {
		return helpers.Error(ctx, 400, "Invalid published year value", err.Error())
	}

	book := &models.Book{
		Author:        ctx.Request().Input("author"),
		Title:         ctx.Request().Input("title"),
		PublishedYear: publishedYear,
		Stock:         stock,
	}

	if err := r.service.UpdateBook(book); err != nil {
		return helpers.Error(ctx, 500, "Failed to update book", err.Error())
	}

	return helpers.Success(ctx, "Book updated successfully", helpers.ToBookResponse(book))
}

func (r *BookController) Destroy(ctx http.Context) http.Response {
	book, err := r.service.GetByIDBook(ctx.Request().Input("id"))
	if err != nil {
		return helpers.Error(ctx, 404, "Book not found", err.Error())
	}
	res, err := r.service.DeleteBook(book)

	if err != nil {
		return helpers.Error(ctx, 500, "Failed to delete book", err.Error())
	}

	return helpers.Success(ctx, "Book deleted successfully", res)
}
