package router

import (
	"api-gateway/handler"

	"api-gateway/middleware"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func Router(redisConn *redis.Client) {
	e := echo.New()

	handler := handler.NewAPIHandler(redisConn)
	m := middleware.NewAuthMiddleware()

	e.POST("/register", handler.Register)

	// student
	e.POST("/document", handler.ProposeDocument, m.IsLogin)

	// staff
	e.PUT("/document", handler.UpdateDocument, m.IsStaff)
	e.POST("/building", handler.NewBuilding, m.IsStaff)
	e.PUT("/building", handler.UpdateBuildingStatus, m.IsStaff)

	e.Logger.Fatal(e.Start(":80"))
}
