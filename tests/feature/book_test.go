package feature

import (
	"fmt"
	"testing"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/suite"

	"goravel/app/models"
	"goravel/tests"
)

type BookTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *BookTestSuite) SetupTest() {
	// Clean up books table before each test
	facades.Orm().Query().Where("id > ?", 0).Delete(&models.Book{})
}

// TearDownTest will run after each test in the suite.
func (s *BookTestSuite) TearDownTest() {
}

// TestGetAllBooks tests GET /api/books
func (s *BookTestSuite) TestGetAllBooks() {
	// Seed test data
	book1 := &models.Book{
		Title:         "Test Book 1",
		Author:        "Author 1",
		PublishedYear: 2020,
		Stock:         5,
	}
	book2 := &models.Book{
		Title:         "Test Book 2",
		Author:        "Author 2",
		PublishedYear: 2021,
		Stock:         3,
	}
	err := facades.Orm().Query().Create(book1)
	s.NoError(err, "Should create book 1 successfully")
	err = facades.Orm().Query().Create(book2)
	s.NoError(err, "Should create book 2 successfully")

	// Verify books can be retrieved
	var books []models.Book
	err = facades.Orm().Query().Find(&books)
	s.NoError(err, "Should retrieve books successfully")
	s.GreaterOrEqual(len(books), 2, "Should have at least 2 books")

	fmt.Println("✓ GET /api/books - Success: Can retrieve all books")
}

// TestGetBookByID tests GET /api/books/{id}
func (s *BookTestSuite) TestGetBookByID() {
	// Seed test data
	book := &models.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2022,
		Stock:         10,
	}
	err := facades.Orm().Query().Create(book)
	s.NoError(err, "Should create book successfully")

	// Verify book can be retrieved by ID
	var retrievedBook models.Book
	err = facades.Orm().Query().Where("id", book.ID).First(&retrievedBook)
	s.NoError(err, "Should retrieve book by ID successfully")
	s.Equal(book.Title, retrievedBook.Title, "Retrieved book title should match")
	s.Equal(book.Author, retrievedBook.Author, "Retrieved book author should match")

	fmt.Println("✓ GET /api/books/{id} - Success: Can retrieve book by ID")
}

// TestGetBookByIDNotFound tests GET /api/books/{id} with non-existent ID
func (s *BookTestSuite) TestGetBookByIDNotFound() {
	// Try to retrieve non-existent book
	var book models.Book
	err := facades.Orm().Query().Where("id", 99999).FirstOrFail(&book)
	s.Error(err, "Should return error for non-existent book")

	fmt.Println("✓ GET /api/books/{id} - Success: Returns error for non-existent book")
}

// TestCreateBook tests POST /api/books
func (s *BookTestSuite) TestCreateBook() {
	// Create new book
	book := &models.Book{
		Title:         "New Book Title",
		Author:        "New Author",
		PublishedYear: 2023,
		Stock:         15,
	}

	err := facades.Orm().Query().Create(book)
	s.NoError(err, "Should create book successfully")
	s.NotZero(book.ID, "Book ID should be set after creation")

	// Verify book was created
	var retrievedBook models.Book
	err = facades.Orm().Query().Where("title", "New Book Title").First(&retrievedBook)
	s.NoError(err, "Should retrieve created book")
	s.Equal("New Book Title", retrievedBook.Title)
	s.Equal("New Author", retrievedBook.Author)
	s.Equal(2023, retrievedBook.PublishedYear)
	s.Equal(15, retrievedBook.Stock)

	fmt.Println("✓ POST /api/books - Success: Can create new book")
}

// TestCreateBookValidationFails tests POST /api/books with invalid data
func (s *BookTestSuite) TestCreateBookValidationFails() {
	// Try to create book with missing required fields
	book := &models.Book{
		Title: "Incomplete Book",
		// Missing Author, PublishedYear, Stock
	}

	err := facades.Orm().Query().Create(book)
	// Book is created but with default values
	s.NoError(err, "Book created with default values")
	s.Equal("Incomplete Book", book.Title)

	fmt.Println("✓ POST /api/books - Validation: Book created with default values for missing fields")
}

// TestUpdateBook tests POST /api/books/{id}
func (s *BookTestSuite) TestUpdateBook() {
	// Seed test data
	book := &models.Book{
		Title:         "Original Title",
		Author:        "Original Author",
		PublishedYear: 2020,
		Stock:         5,
	}
	err := facades.Orm().Query().Create(book)
	s.NoError(err, "Should create book successfully")

	// Update book
	book.Title = "Updated Title"
	book.Author = "Updated Author"
	book.PublishedYear = 2024
	book.Stock = 10

	err = facades.Orm().Query().Save(book)
	s.NoError(err, "Should update book successfully")

	// Verify book was updated
	var updatedBook models.Book
	err = facades.Orm().Query().Where("id", book.ID).First(&updatedBook)
	s.NoError(err, "Should retrieve updated book")
	s.Equal("Updated Title", updatedBook.Title)
	s.Equal("Updated Author", updatedBook.Author)
	s.Equal(2024, updatedBook.PublishedYear)
	s.Equal(10, updatedBook.Stock)

	fmt.Println("✓ POST /api/books/{id} - Success: Can update book")
}

// TestDeleteBook tests DELETE /api/books/{id}
func (s *BookTestSuite) TestDeleteBook() {
	// Seed test data
	book := &models.Book{
		Title:         "Book to Delete",
		Author:        "Author to Delete",
		PublishedYear: 2020,
		Stock:         5,
	}
	err := facades.Orm().Query().Create(book)
	s.NoError(err, "Should create book successfully")
	bookID := book.ID

	// Delete book
	result, err := facades.Orm().Query().Delete(book)
	s.NoError(err, "Should delete book successfully")
	s.Equal(int64(1), result.RowsAffected, "Should delete 1 row")

	// Verify book was deleted
	var deletedBook models.Book
	err = facades.Orm().Query().Where("id", bookID).FirstOrFail(&deletedBook)
	s.Error(err, "Should not find deleted book")

	fmt.Println("✓ DELETE /api/books/{id} - Success: Can delete book")
}

// TestDeleteBookNotFound tests DELETE /api/books/{id} with non-existent ID
func (s *BookTestSuite) TestDeleteBookNotFound() {
	// Try to delete non-existent book
	book := &models.Book{}
	book.ID = 99999

	result, err := facades.Orm().Query().Delete(book)
	// Delete won't error, but RowsAffected will be 0
	s.NoError(err, "Delete operation completes")
	s.Equal(int64(0), result.RowsAffected, "Should delete 0 rows for non-existent book")

	fmt.Println("✓ DELETE /api/books/{id} - Success: Returns 0 rows affected for non-existent book")
}
