package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250814032821CreateBorrowingsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250814032821CreateBorrowingsTable) Signature() string {
	return "20250814032821_create_borrowings_table"
}

// Up Run the migrations.
func (r *M20250814032821CreateBorrowingsTable) Up() error {
	if !facades.Schema().HasTable("borrowings") {
		return facades.Schema().Create("borrowings", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("user_id")
			table.Foreign("user_id").References("id").On("users").CascadeOnUpdate().CascadeOnDelete()
			table.UnsignedBigInteger("book_id")
			table.Foreign("book_id").References("id").On("books").CascadeOnUpdate().CascadeOnDelete()
			table.Date("borrow_date")
			table.Date("return_date").Nullable()
			table.Enum("status", []any{"borrowed", "returned"})
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250814032821CreateBorrowingsTable) Down() error {
	return facades.Schema().DropIfExists("borrowings")
}
