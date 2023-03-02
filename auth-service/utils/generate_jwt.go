package utils

import (
	"auth-service/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user *models.User) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    user.ID,
		"name":      user.Name,
		"role":      user.Role,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenStr, err := claims.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
