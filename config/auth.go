package config

import (
	"errors"
	"github.com/afanasyevadina/go-test/models"
	"github.com/afanasyevadina/go-test/services"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

var CurrentUser *models.User

func Authenticate(r *http.Request) error {
	if err := tryToken(r.FormValue("token")); err == nil {
		return nil
	}
	token, _ := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	return tryToken(token)
}

func tryToken(token string) error {
	claims, err := services.GetJwtService().ParseToken(token)
	subject, err := claims.GetSubject()
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(subject)
	if err != nil {
		return err
	}
	res := DB.First(&CurrentUser, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}
