package controllers

import (
	"errors"
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/dto"
	"github.com/afanasyevadina/go-test/models"
	"github.com/afanasyevadina/go-test/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func NewAuthController() AuthController {
	return AuthController{
		jwtService: services.GetJwtService(),
	}
}

type AuthController struct {
	jwtService *services.JwtService
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := dto.LoginRequest{}
	dto.FromRequest(r, &loginRequest)
	if loginRequest.Email == "" || loginRequest.Password == "" {
		dto.ToJsonResponse(w, dto.ValidationErrorResponse{
			Errors: map[string]string{"email": "required", "password": "required"},
		}, http.StatusUnprocessableEntity)
		return
	}
	user := models.User{}
	res := config.DB.Where("email = ?", loginRequest.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil || errors.Is(res.Error, gorm.ErrRecordNotFound) {
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
	if registerRequest.Email == "" || registerRequest.Password == "" {
		dto.ToJsonResponse(w, dto.ValidationErrorResponse{
			Errors: map[string]string{"email": "required", "password": "required"},
		}, http.StatusUnprocessableEntity)
		return
	}
	user := models.User{Email: registerRequest.Email, Name: registerRequest.Name}
	res := config.DB.Where("email = ?", registerRequest.Email).First(&user)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		dto.ToJsonResponse(w, dto.ValidationErrorResponse{
			Errors: map[string]string{"email": "required", "password": "required"},
		}, http.StatusUnprocessableEntity)
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	config.DB.Save(&user)
	token, err := c.jwtService.CreateToken(user.ID)
	if err != nil {
		dto.RespondWith400(w)
		return
	}
	dto.ToJsonResponse(w, dto.TokenResponse{
		Token: token,
	}, http.StatusOK)
}
