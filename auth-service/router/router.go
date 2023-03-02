package router

import (
	"auth-service/handler"
	"auth-service/repository"
	usecase "auth-service/use-case"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(mysqlConn *gorm.DB) {
	e := echo.New()

	repo := repository.NewAuthRepo(mysqlConn)
	useCase := usecase.NewAuthUseCase(repo)
	handler := handler.NewAuthHandler(useCase)

	e.POST("/auth", handler.Auth)

	e.Logger.Fatal(e.Start(":80"))
}
