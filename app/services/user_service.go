package services

import (
	"errors"
	"goravel/app/models"
	"goravel/app/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/facades"
)

type UserService interface {
	Login(email, password string) (*models.User, string, error)
	Logout(token string) error
	RegisterUser(user *models.User) error
	GetAllUser() ([]models.User, error)
	GetByIDUser(id any) (*models.User, error)
	UpdateUser(user *models.User, id int) error
	DeleteUser(user *models.User) (int64, error)
	ValidateEmailUnique(email string, excludeID int) error
}

type userService struct {
	repo repositories.UserRepository
}

var ErrEmailExists = errors.New("email already exists")

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUser() ([]models.User, error) {
	return s.repo.FindAllUser()
}

func (s *userService) GetByIDUser(id any) (*models.User, error) {
	return s.repo.FindByIDUser(id)
}

func (s *userService) RegisterUser(user *models.User) error {
	if err := s.ValidateEmailUnique(user.Email, 0); err != nil {
		return err
	}

	hashedPassword, err := facades.Hash().Make(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return s.repo.RegisterUser(user)
}

func (s *userService) UpdateUser(user *models.User, id int) error {
	if err := s.ValidateEmailUnique(user.Email, id); err != nil {
		return err
	}

	hashedPassword, err := facades.Hash().Make(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(user *models.User) (int64, error) {
	return s.repo.DeleteUser(user)
}

func (s *userService) ValidateEmailUnique(email string, id int) error {
	isUnique, err := s.repo.ExcludeEmailByID(email, id)
	if err != nil {
		return err
	}
	if !isUnique {
		return ErrEmailExists
	}

	return nil
}

func (s *userService) Login(email, password string) (*models.User, string, error) {
	user, err := s.repo.LoginUser(email, password)
	if err != nil {
		return nil, "", err
	}

	// Generate JWT token manually
	jwtSecret := facades.Config().GetString("jwt.secret")
	if jwtSecret == "" {
		return nil, "", errors.New("JWT secret not configured")
	}

	ttl := facades.Config().GetInt("jwt.ttl", 60)

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Duration(ttl) * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}

func (s *userService) Logout(token string) error {
	return s.repo.Logout(token)
}
