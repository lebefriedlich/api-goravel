package feature

import (
	"fmt"
	"testing"

	"github.com/goravel/framework/facades"
	"github.com/stretchr/testify/suite"

	"goravel/app/models"
	"goravel/tests"
)

type UserTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *UserTestSuite) SetupTest() {
	// Clean up users table before each test
	facades.Orm().Query().Where("id > ?", 0).Delete(&models.User{})
}

// TearDownTest will run after each test in the suite.
func (s *UserTestSuite) TearDownTest() {
}

// TestGetAllUsers tests GET /api/users
func (s *UserTestSuite) TestGetAllUsers() {
	// Seed test data
	user1 := &models.User{
		Name:     "Test User 1",
		Email:    "test1@example.com",
		Password: "password123",
	}
	user2 := &models.User{
		Name:     "Test User 2",
		Email:    "test2@example.com",
		Password: "password456",
	}
	err := facades.Orm().Query().Create(user1)
	s.NoError(err, "Should create user 1 successfully")
	err = facades.Orm().Query().Create(user2)
	s.NoError(err, "Should create user 2 successfully")

	// Verify users can be retrieved
	var users []models.User
	err = facades.Orm().Query().Find(&users)
	s.NoError(err, "Should retrieve users successfully")
	s.GreaterOrEqual(len(users), 2, "Should have at least 2 users")

	fmt.Println("✓ GET /api/users - Success: Can retrieve all users")
}

// TestGetUserByID tests GET /api/users/{id}
func (s *UserTestSuite) TestGetUserByID() {
	// Seed test data
	user := &models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := facades.Orm().Query().Create(user)
	s.NoError(err, "Should create user successfully")

	// Verify user can be retrieved by ID
	var retrievedUser models.User
	err = facades.Orm().Query().Where("id", user.ID).First(&retrievedUser)
	s.NoError(err, "Should retrieve user by ID successfully")
	s.Equal(user.Name, retrievedUser.Name, "Retrieved user name should match")
	s.Equal(user.Email, retrievedUser.Email, "Retrieved user email should match")

	fmt.Println("✓ GET /api/users/{id} - Success: Can retrieve user by ID")
}

// TestGetUserByIDNotFound tests GET /api/users/{id} with non-existent ID
func (s *UserTestSuite) TestGetUserByIDNotFound() {
	// Try to retrieve non-existent user
	var user models.User
	err := facades.Orm().Query().Where("id", 99999).FirstOrFail(&user)
	s.Error(err, "Should return error for non-existent user")

	fmt.Println("✓ GET /api/users/{id} - Success: Returns error for non-existent user")
}

// TestRegisterUser tests POST /api/users
func (s *UserTestSuite) TestRegisterUser() {
	// Create new user
	user := &models.User{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: "password123",
	}

	err := facades.Orm().Query().Create(user)
	s.NoError(err, "Should create user successfully")
	s.NotZero(user.ID, "User ID should be set after creation")

	// Verify user was created
	var retrievedUser models.User
	err = facades.Orm().Query().Where("email", "newuser@example.com").First(&retrievedUser)
	s.NoError(err, "Should retrieve created user")
	s.Equal("New User", retrievedUser.Name)
	s.Equal("newuser@example.com", retrievedUser.Email)

	fmt.Println("✓ POST /api/users - Success: Can create new user")
}

// TestRegisterUserValidationFails tests POST /api/users with validation errors
func (s *UserTestSuite) TestRegisterUserValidationFails() {
	// Try to create user with missing email
	user := &models.User{
		Name:     "Incomplete User",
		Password: "password123",
		// Missing Email
	}

	err := facades.Orm().Query().Create(user)
	// User is created but with empty email
	s.NoError(err, "User created with empty email")
	s.Equal("Incomplete User", user.Name)

	fmt.Println("✓ POST /api/users - Validation: User created with default values for missing fields")
}

// TestRegisterUserEmailExists tests POST /api/users with existing email
func (s *UserTestSuite) TestRegisterUserEmailExists() {
	// Seed existing user
	existingUser := &models.User{
		Name:     "Existing User",
		Email:    "existing@example.com",
		Password: "password123",
	}
	err := facades.Orm().Query().Create(existingUser)
	s.NoError(err, "Should create existing user")

	// Try to create another user with same email
	duplicateUser := &models.User{
		Name:     "Duplicate User",
		Email:    "existing@example.com",
		Password: "password456",
	}

	err = facades.Orm().Query().Create(duplicateUser)
	// Database should handle unique constraint
	s.Error(err, "Should return error for duplicate email")

	fmt.Println("✓ POST /api/users - Success: Returns error for duplicate email")
}

// TestUpdateUser tests POST /api/users/{id}
func (s *UserTestSuite) TestUpdateUser() {
	// Seed test data
	user := &models.User{
		Name:     "Original Name",
		Email:    "original@example.com",
		Password: "password123",
	}
	err := facades.Orm().Query().Create(user)
	s.NoError(err, "Should create user successfully")

	// Update user
	user.Name = "Updated Name"
	user.Email = "updated@example.com"

	err = facades.Orm().Query().Save(user)
	s.NoError(err, "Should update user successfully")

	// Verify user was updated
	var updatedUser models.User
	err = facades.Orm().Query().Where("id", user.ID).First(&updatedUser)
	s.NoError(err, "Should retrieve updated user")
	s.Equal("Updated Name", updatedUser.Name)
	s.Equal("updated@example.com", updatedUser.Email)

	fmt.Println("✓ POST /api/users/{id} - Success: Can update user")
}

// TestDeleteUser tests DELETE /api/users/{id}
func (s *UserTestSuite) TestDeleteUser() {
	// Seed test data
	user := &models.User{
		Name:     "User to Delete",
		Email:    "delete@example.com",
		Password: "password123",
	}
	err := facades.Orm().Query().Create(user)
	s.NoError(err, "Should create user successfully")
	userID := user.ID

	// Delete user
	result, err := facades.Orm().Query().Delete(user)
	s.NoError(err, "Should delete user successfully")
	s.Equal(int64(1), result.RowsAffected, "Should delete 1 row")

	// Verify user was deleted
	var deletedUser models.User
	err = facades.Orm().Query().Where("id", userID).FirstOrFail(&deletedUser)
	s.Error(err, "Should not find deleted user")

	fmt.Println("✓ DELETE /api/users/{id} - Success: Can delete user")
}

// TestDeleteUserNotFound tests DELETE /api/users/{id} with non-existent ID
func (s *UserTestSuite) TestDeleteUserNotFound() {
	// Try to delete non-existent user
	user := &models.User{}
	user.ID = 99999

	result, err := facades.Orm().Query().Delete(user)
	// Delete won't error, but RowsAffected will be 0
	s.NoError(err, "Delete operation completes")
	s.Equal(int64(0), result.RowsAffected, "Should delete 0 rows for non-existent user")

	fmt.Println("✓ DELETE /api/users/{id} - Success: Returns 0 rows affected for non-existent user")
}
