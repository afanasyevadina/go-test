package controllers

import (
	"encoding/json"
	"errors"
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/models"
	"github.com/afanasyevadina/go-test/services"
	"github.com/afanasyevadina/go-test/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func NewAuthController() AuthController {
	return AuthController{
		jwtService: services.JwtService{},
	}
}

type AuthController struct {
	jwtService services.JwtService
}

type Credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	credentials := Credentials{}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if credentials.Email == "" || credentials.Password == "" {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
		}, http.StatusUnprocessableEntity)
		return
	}
	user := models.User{}
	res := config.DB.Where("email = ?", credentials.Email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		}, http.StatusUnauthorized)
		return
	}
	token, err := c.jwtService.CreateToken(user.ID)
	if err != nil {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
		return
	}
	util.JsonResponse(w, Token{token}, http.StatusOK)
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	credentials := Credentials{}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if credentials.Email == "" || credentials.Password == "" {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
		}, http.StatusUnprocessableEntity)
		return
	}
	user := models.User{Email: credentials.Email}
	res := config.DB.Where("email = ?", credentials.Email).First(&user)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
		}, http.StatusUnprocessableEntity)
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	config.DB.Save(&user)
	token, err := c.jwtService.CreateToken(user.ID)
	if err != nil {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
		return
	}
	util.JsonResponse(w, Token{token}, http.StatusOK)
}
