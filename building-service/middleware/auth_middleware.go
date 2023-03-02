package middleware

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	IsLogin(next echo.HandlerFunc) echo.HandlerFunc
	IsStaff(next echo.HandlerFunc) echo.HandlerFunc
}
