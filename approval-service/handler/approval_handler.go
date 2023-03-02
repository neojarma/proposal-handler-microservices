package handler

import (
	"github.com/labstack/echo/v4"
)

type ApprovalHandler interface {
	NewDocument(payload string)
	UpdateProgress(payload string)
	GetDocument(ctx echo.Context) error
}
