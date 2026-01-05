package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250814025201CreateBooksTable struct{}

// Signature The unique signature for the migration.
func (r *M20250814025201CreateBooksTable) Signature() string {
	return "20250814025201_create_books_table"
}

// Up Run the migrations.
func (r *M20250814025201CreateBooksTable) Up() error {
	if !facades.Schema().HasTable("books") {
		return facades.Schema().Create("books", func(table schema.Blueprint) {
			table.ID()
			table.String("title", 255)
			table.String("author", 100)
			table.Integer("published_year")
			table.Integer("stock")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250814025201CreateBooksTable) Down() error {
	return facades.Schema().DropIfExists("books")
}
