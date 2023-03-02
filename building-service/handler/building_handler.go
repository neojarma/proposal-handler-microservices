package handler

import "github.com/labstack/echo/v4"

type BuildingHandler interface {
	GetBuildings(ctx echo.Context) error
	NewBuilding(payload string)
	UpdateBuildingStatus(payload string)
}
