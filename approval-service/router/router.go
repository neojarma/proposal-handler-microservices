package router

import (
	"approval-service/handler"
	"approval-service/middleware"
	"approval-service/repository"
	usecase "approval-service/use-case"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	broker "approval-service/redis"
)

func Router(db *gorm.DB, redisConn *redis.Client) {
	e := echo.New()

	repo := repository.NewApprovalRepo(db)
	useCase := usecase.NewApprovalUseCase(repo)
	handler := handler.NewApprovalHandler(useCase, redisConn)
	m := middleware.NewAuthMiddleware()

	go broker.ListenRedisEvent(redisConn, os.Getenv("NEW_DOCUMENT_CHANNEL"), handler.NewDocument)
	go broker.ListenRedisEvent(redisConn, os.Getenv("UPDATE_DOCUMENT_CHANNEL"), handler.UpdateProgress)

	e.GET("/document", handler.GetDocument, m.IsLogin)

	e.Logger.Fatal(e.Start(":80"))
}
