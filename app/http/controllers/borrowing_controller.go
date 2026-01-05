package controllers

import (
	"goravel/app/helpers"
	"goravel/app/models"
	"goravel/app/repositories"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/http"
)

type BorrowingController struct {
	service services.BorrowingService
}

func NewBorrowingController() *BorrowingController {
	repo := repositories.NewBorrowingRepository()
	service := services.NewBorrowingService(repo)
	return &BorrowingController{service: service}
}

func (r *BorrowingController) Index(ctx http.Context) http.Response {
	borrowings, err := r.service.GetAllBorrowings()
	if err != nil {
		return helpers.Error(ctx, 500, "Failed to fetch borrowings", err.Error)
	}
	borrowingResponses := helpers.ToBorrowingResponseList(borrowings)
	return helpers.Success(ctx, "Borrowings retrieved successfully", borrowingResponses)
}

func (r *BorrowingController) Borrow(ctx http.Context) http.Response {
	borrowing := &models.Borrowing{}
	err := ctx.Request().Bind(borrowing)
	if err != nil {
		return helpers.Error(ctx, 400, "Failed to bind borrowing data", err.Error())
	}

	userID := ctx.Request().Input("user_id")
	bookID := ctx.Request().Input("book_id")

	err = r.service.BorrowingUser(borrowing, userID, bookID)
	if err != nil {
		return helpers.Error(ctx, 500, "Failed to borrow book", err.Error())
	}

	return helpers.Success(ctx, "Book borrowed successfully", helpers.ToBorrowingResponse(borrowing))
}

func (r *BorrowingController) Return(ctx http.Context) http.Response {
	borrowing := &models.Borrowing{}
	err := ctx.Request().Bind(borrowing)
	if err != nil {
		return helpers.Error(ctx, 400, "Failed to bind borrowing data", err.Error())
	}

	userID := ctx.Request().Input("user_id")
	bookID := ctx.Request().Input("book_id")

	err = r.service.ReturnUserBorrowing(borrowing, userID, bookID)
	if err != nil {
		return helpers.Error(ctx, 500, "Failed to return book", err.Error())
	}

	return helpers.Success(ctx, "Book returned successfully", helpers.ToBorrowingResponse(borrowing))
}

func (r *BorrowingController) FindByUserID(ctx http.Context) http.Response {
	userID := ctx.Request().Input("user_id")

	borrowing, err := r.service.FindByUserIDBorrowing(userID)
	if err != nil {
		return helpers.Error(ctx, 500, "Failed to fetch borrowing data", err.Error())
	}

	return helpers.Success(ctx, "Borrowing data retrieved successfully", helpers.ToBorrowingResponse(borrowing))
}
