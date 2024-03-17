package dto

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/models"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func UserResponseFromModel(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (req UserUpdateRequest) ToCurrentUser() {
	(*config.CurrentUser).Email = req.Email
	(*config.CurrentUser).Name = req.Name
}
