package controllers

import (
	"goravel/app/helpers"
	"goravel/app/models"
	"goravel/app/repositories"
	"goravel/app/services"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	service services.UserService
}

func NewUserController() *UserController {
	repo := repositories.NewUserRepository()
	service := services.NewUserService(repo)
	return &UserController{service: service}
}

func (r *UserController) Login(ctx http.Context) http.Response {
	validation, err := ctx.Request().Validate(map[string]string{
		"email":    "required|string|email|max_len:255",
		"password": "required|string|min_len:8",
	})

	if err != nil {
		return helpers.Error(ctx, 500, "Validation setup failed", err.Error())
	}

	if validation.Fails() {
		return helpers.Error(ctx, 400, "Validation failed", validation.Errors().All())
	}

	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	user, token, err := r.service.Login(email, password)
	if err != nil {
		return helpers.Error(ctx, 401, "Login failed", err.Error())
	}

	return helpers.Success(ctx, "Login successful", map[string]any{
		"user":  helpers.ToUserResponse(user),
		"token": token,
	})
}

func (r *UserController) Logout(ctx http.Context) http.Response {
	err := r.service.Logout(ctx.Request().Input("token"))
	if err != nil {
		return helpers.Error(ctx, 500, "Logout failed", err.Error())
	}

	return helpers.Success(ctx, "Logout successful", nil)
}

func (r *UserController) Index(ctx http.Context) http.Response {
	users, err := r.service.GetAllUser()
	if err != nil {
		return helpers.Error(ctx, 500, "Failed to fetch users", err.Error())
	}

	userResponses := helpers.ToUserResponseList(users)

	return helpers.Success(ctx, "Users retrieved successfully", userResponses)
}

func (r *UserController) Register(ctx http.Context) http.Response {
	validation, err := ctx.Request().Validate(map[string]string{
		"name":     "required|string|max_len:255",
		"email":    "required|string|email|max_len:255",
		"password": "required|string|min_len:8",
	})

	if err != nil {
		return helpers.Error(ctx, 500, "Validation setup failed", err.Error())
	}

	if validation.Fails() {
		return helpers.Error(ctx, 400, "Validation failed", validation.Errors().All())
	}

	user := &models.User{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("email"),
		Password: ctx.Request().Input("password"),
	}

	if err := r.service.RegisterUser(user); err != nil {
		if err.Error() == "email already exists" {
			return helpers.Error(ctx, 400, "Email already exists", nil)
		}
		return helpers.Error(ctx, 500, "Failed to create user", err.Error())
	}

	return helpers.Success(ctx, "User created successfully", helpers.ToUserResponse(user))
}

func (r *UserController) Show(ctx http.Context) http.Response {
	user, err := r.service.GetByIDUser(ctx.Request().Input("id"))
	if err != nil {
		return helpers.Error(ctx, 404, "User not found", err.Error())
	}

	return helpers.Success(ctx, "User retrieved successfully", helpers.ToUserResponse(user))
}

func (r *UserController) Update(ctx http.Context) http.Response {
	validation, err := ctx.Request().Validate(map[string]string{
		"name":     "string|max_len:255",
		"email":    "string|email|max_len:255",
		"password": "string|min_len:8",
	})

	if err != nil {
		return helpers.Error(ctx, 500, "Validation setup failed", err.Error())
	}

	if validation.Fails() {
		return helpers.Error(ctx, 400, "Validation failed", validation.Errors().All())
	}

	user, err := r.service.GetByIDUser(ctx.Request().Input("id"))
	if err != nil {
		return helpers.Error(ctx, 404, "User not found", err.Error())
	}

	idStr := ctx.Request().Input("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.Error(ctx, 400, "Invalid user ID", err.Error())
	}

	if err := r.service.UpdateUser(user, id); err != nil {
		return helpers.Error(ctx, 500, "Failed to update user", err.Error())
	}

	return helpers.Success(ctx, "User updated successfully", helpers.ToUserResponse(user))
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	user, err := r.service.GetByIDUser(ctx.Request().Input("id"))
	if err != nil {
		return helpers.Error(ctx, 404, "User not found", err.Error())
	}
	res, err := r.service.DeleteUser(user)

	if err != nil {
		return helpers.Error(ctx, 500, "Failed to delete user", err.Error())
	}

	return helpers.Success(ctx, "User deleted successfully", res)
}
