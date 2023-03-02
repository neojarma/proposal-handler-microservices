package handler

import "github.com/labstack/echo/v4"

type APIHandler interface {
	Register(ctx echo.Context) error
	ProposeDocument(ctx echo.Context) error
	UpdateDocument(ctx echo.Context) error
	NewBuilding(ctx echo.Context) error
	UpdateBuildingStatus(ctx echo.Context) error
}
