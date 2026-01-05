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
		return helpers.Error(ctx, 500, "Gagal mengambil data peminjaman", err.Error)
	}
	borrowingResponses := helpers.ToBorrowingResponseList(borrowings)
	return helpers.Success(ctx, "Pengambilan data peminjaman berhasil", borrowingResponses)
}

func (r *BorrowingController) Borrow(ctx http.Context) http.Response {
	borrowing := &models.Borrowing{}
	err := ctx.Request().Bind(borrowing)
	if err != nil {
		return helpers.Error(ctx, 400, "Gagal mengikat data peminjaman", err.Error())
	}

	userID := ctx.Request().Input("user_id")
	bookID := ctx.Request().Input("book_id")

	err = r.service.BorrowingUser(borrowing, userID, bookID)
	if err != nil {
		return helpers.Error(ctx, 500, "Gagal meminjam buku", err.Error())
	}

	return helpers.Success(ctx, "Buku berhasil dipinjam", helpers.ToBorrowingResponse(borrowing))
}

func (r *BorrowingController) Return(ctx http.Context) http.Response {
	borrowing := &models.Borrowing{}
	err := ctx.Request().Bind(borrowing)
	if err != nil {
		return helpers.Error(ctx, 400, "Gagal mengikat data peminjaman", err.Error())
	}

	userID := ctx.Request().Input("user_id")
	bookID := ctx.Request().Input("book_id")

	err = r.service.ReturnUserBorrowing(borrowing, userID, bookID)
	if err != nil {
		return helpers.Error(ctx, 500, "Gagal mengembalikan buku", err.Error())
	}

	return helpers.Success(ctx, "Buku berhasil dikembalikan", helpers.ToBorrowingResponse(borrowing))
}

func (r *BorrowingController) FindByUserID(ctx http.Context) http.Response {
	userID := ctx.Request().Input("user_id")

	borrowing, err := r.service.FindByUserIDBorrowing(userID)
	if err != nil {
		return helpers.Error(ctx, 500, "Gagal mengambil data peminjaman", err.Error())
	}

	return helpers.Success(ctx, "Pengambilan data peminjaman berhasil", helpers.ToBorrowingResponse(borrowing))
}
