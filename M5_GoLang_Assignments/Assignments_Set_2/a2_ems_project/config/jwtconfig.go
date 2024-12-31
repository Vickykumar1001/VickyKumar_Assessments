package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = "mysecretkey"

func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}
