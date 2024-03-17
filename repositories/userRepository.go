package repositories

import (
	"errors"
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type UserRepository struct {
}

func (repo UserRepository) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	res := config.DB.Where("email = ?", email).First(&user)
	return user, res.Error
}

func (repo UserRepository) LoginByEmail(email string, password string) (models.User, error) {
	user, err := repo.GetByEmail(email)
	if err != nil {
		return user, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, err
	}
	return user, nil
}

func (repo UserRepository) Create(user models.User) (models.User, error) {
	_, err := repo.GetByEmail(user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("email is already taken")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	config.DB.Save(&user)
	return user, nil
}

func (repo UserRepository) Update(user models.User) models.User {
	config.DB.Save(&user)
	return user
}
