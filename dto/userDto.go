package dto

import "github.com/afanasyevadina/go-test/models"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserResponseFromModel(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
