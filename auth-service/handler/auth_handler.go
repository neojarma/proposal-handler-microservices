package handler

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Auth(ctx echo.Context) error
}
