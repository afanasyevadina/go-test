package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type JwtService struct {
	secret []byte
}

func GetJwtService() *JwtService {
	return &JwtService{[]byte("secret-key")}
}

func (s *JwtService) CreateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": strconv.Itoa(int(userId)),
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}
