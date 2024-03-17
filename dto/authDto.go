package dto

import "github.com/afanasyevadina/go-test/models"

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req RegisterRequest) ToModel() models.User {
	return models.User{Email: req.Email, Name: req.Name}
}
