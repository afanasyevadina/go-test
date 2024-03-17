package controllers

import (
	"github.com/afanasyevadina/go-test/dto"
	"github.com/afanasyevadina/go-test/repositories"
	"github.com/afanasyevadina/go-test/services"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewAuthController() AuthController {
	return AuthController{
		jwtService:     services.NewJwtService(),
		userRepository: repositories.NewUserRepository(),
		validator:      validator.New(validator.WithRequiredStructEnabled()),
	}
}

type AuthController struct {
	jwtService     *services.JwtService
	userRepository *repositories.UserRepository
	validator      *validator.Validate
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := dto.LoginRequest{}
	dto.FromRequest(r, &loginRequest)
	if err := c.validator.Struct(loginRequest); err != nil {
		dto.ToJsonResponse(w, dto.ResponseFromValidator(err.(validator.ValidationErrors)), http.StatusUnprocessableEntity)
		return
	}
	user, err := c.userRepository.LoginByEmail(loginRequest.Email, loginRequest.Password)
	if err != nil {
		dto.RespondWith401(w)
		return
	}
	token, err := c.jwtService.CreateToken(user.ID)
	if err != nil {
		dto.RespondWith400(w)
		return
	}
	dto.ToJsonResponse(w, dto.TokenResponse{
		Token: token,
	}, http.StatusOK)
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	registerRequest := dto.RegisterRequest{}
	dto.FromRequest(r, &registerRequest)
	err := c.validator.Struct(registerRequest)
	if err := c.validator.Struct(registerRequest); err != nil {
		dto.ToJsonResponse(w, dto.ResponseFromValidator(err.(validator.ValidationErrors)), http.StatusUnprocessableEntity)
		return
	}
	user, err := c.userRepository.Create(registerRequest.ToModel())
	if err != nil {
		dto.ToJsonResponse(w, dto.ValidationErrorResponse{Errors: map[string]string{"email": err.Error()}}, http.StatusUnprocessableEntity)
		return
	}
	token, err := c.jwtService.CreateToken(user.ID)
	if err != nil {
		dto.RespondWith400(w)
		return
	}
	dto.ToJsonResponse(w, dto.TokenResponse{
		Token: token,
	}, http.StatusOK)
}
