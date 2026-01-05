package models

import (
	"github.com/goravel/framework/database/orm"
)

type Borrowing struct {
	orm.Model
	UserID     uint
	BookID     uint
	BorrowDate string
	ReturnDate string
	Status     string
}
