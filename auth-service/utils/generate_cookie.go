package utils

import (
	"net/http"
	"time"
)

func GenerateCookie(name, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(15 * time.Minute)

	cookie.Name = name
	cookie.Value = value

	return cookie
}
