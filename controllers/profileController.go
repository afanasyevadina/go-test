package controllers

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/dto"
	"github.com/afanasyevadina/go-test/repositories"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewProfileController() ProfileController {
	return ProfileController{
		userRepository: repositories.NewUserRepository(),
		validator:      validator.New(validator.WithRequiredStructEnabled()),
	}
}

type ProfileController struct {
	userRepository *repositories.UserRepository
	validator      *validator.Validate
}

func (c *ProfileController) Show(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c.Update(w, r)
	} else {
		dto.ToJsonResponse(w, dto.UserResponseFromModel(*config.CurrentUser), http.StatusOK)
	}
}

func (c *ProfileController) Update(w http.ResponseWriter, r *http.Request) {
	userRequest := dto.UserUpdateRequest{}
	dto.FromRequest(r, &userRequest)
	if err := c.validator.Struct(userRequest); err != nil {
		dto.ToJsonResponse(w, dto.ResponseFromValidator(err.(validator.ValidationErrors)), http.StatusUnprocessableEntity)
		return
	}
	userRequest.ToCurrentUser()
	c.userRepository.Update(*config.CurrentUser)
	dto.ToStatusResponse(w, http.StatusOK)
}
