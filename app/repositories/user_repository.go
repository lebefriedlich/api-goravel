package repositories

import (
	"errors"
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UserRepository interface {
	LoginUser(email, password string) (*models.User, error)
	FindAllUser() ([]models.User, error)
	FindByIDUser(id any) (*models.User, error)
	ExcludeEmailByID(email string, id int) (bool, error)
	RegisterUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) (int64, error)
	Logout(token string) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email, password string) (*models.User, error) {
	var user models.User
	err := facades.Orm().Query().Where("email", email).First(&user)
	if err != nil {
		return nil, err
	}

	if !facades.Hash().Check(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (r *userRepository) FindAllUser() ([]models.User, error) {
	var users []models.User
	err := facades.Orm().Query().Find(&users)
	return users, err
}

func (r *userRepository) FindByIDUser(id any) (*models.User, error) {
	var user models.User
	err := facades.Orm().Query().Where("id", id).FirstOrFail(&user)
	return &user, err
}

func (r *userRepository) ExcludeEmailByID(email string, id int) (bool, error) {
	query := facades.Orm().Query().Model(&models.User{}).Where("email = ?", email)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	count, err := query.Count()
	return count == 0, err
}

func (r *userRepository) RegisterUser(user *models.User) error {
	return facades.Orm().Query().Create(user)
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return facades.Orm().Query().Save(user)
}

func (r *userRepository) DeleteUser(user *models.User) (int64, error) {
	res, err := facades.Orm().Query().Delete(user)
	return res.RowsAffected, err
}

func (r *userRepository) Logout(token string) error {
	_, err := facades.Orm().Query().Where("token", token).Update("token", "")
	return err
}
