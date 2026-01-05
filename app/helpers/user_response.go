package helpers

import (
	"goravel/app/models"
)

type UserResponse map[string]any

func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}
}

func ToUserResponseList(users []models.User) []UserResponse {
	var response []UserResponse
	for _, user := range users {
		response = append(response, ToUserResponse(&user))
	}
	return response
}
