package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

// verify the token and return role
func VerifyJWT(token string) (string, bool) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("Error verifying token:", err)
		return "", false
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		fmt.Println("Invalid token")
		return "", false
	}

	return claims["role"].(string), true
}
