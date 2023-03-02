package utils

import "os"

func CheckTokenService(token string) bool {
	return token == os.Getenv("TOKEN_SERVICE")
}
