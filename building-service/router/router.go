package router

import (
	"building-service/handler"
	"building-service/middleware"
	broker "building-service/redis"
	"building-service/repository"
	usecase "building-service/use-case"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(mysqlConn *gorm.DB, rediConn *redis.Client) {
	e := echo.New()

	repo := repository.NewBuildingRepo(mysqlConn)
	useCase := usecase.NewBuildingUseCase(repo)
	handler := handler.NewBuildingHandler(useCase)
	m := middleware.NewAuthMiddleware()

	go broker.ListenRedisEvent(rediConn, os.Getenv("NEW_BUILDING_CHANNEL"), handler.NewBuilding)
	go broker.ListenRedisEvent(rediConn, os.Getenv("UPDATE_BUILDING_CHANNEL"), handler.UpdateBuildingStatus)

	e.GET("/buildings", handler.GetBuildings, m.IsLogin)

	e.Logger.Fatal(e.Start(":80"))
}
