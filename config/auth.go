package config

import (
	"errors"
	"github.com/afanasyevadina/go-test/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var CurrentUser *models.User

func Authenticate(r *http.Request) error {
	token := r.FormValue("token")
	id, err := strconv.Atoi(token)
	if err != nil {
		return err
	}
	res := DB.First(&CurrentUser, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	return nil
}
