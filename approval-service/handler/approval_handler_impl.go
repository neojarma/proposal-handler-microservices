package handler

import (
	model "approval-service/models"
	useCase "approval-service/use-case"
	"approval-service/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type ApprovalHandlerImpl struct {
	UseCase useCase.ApprovalUseCase
	Redis   *redis.Client
}

func NewApprovalHandler(useCase useCase.ApprovalUseCase, redisConn *redis.Client) ApprovalHandler {
	return &ApprovalHandlerImpl{
		UseCase: useCase,
		Redis:   redisConn,
	}
}

func (handler *ApprovalHandlerImpl) NewDocument(payload string) {

	payloadService := new(model.PayloadServices)
	err := json.Unmarshal([]byte(payload), payloadService)
	if err != nil {
		log.Println("error while unmarshal payload, err:", err)
		return
	}

	if ok := utils.CheckTokenService(payloadService.SecretToken); !ok {
		log.Println("unauthorized access")
		return
	}

	err = handler.UseCase.NewDocument(payloadService.DataDocument)
	if err != nil {
		log.Println("error create new document, err: ", err)
		return
	}

}

func (handler *ApprovalHandlerImpl) UpdateProgress(payload string) {

	payloadService := new(model.PayloadServices)
	err := json.Unmarshal([]byte(payload), payloadService)
	if err != nil {
		log.Println("error while unmarshal payload, err:", err)
	}

	if ok := utils.CheckTokenService(payloadService.SecretToken); !ok {
		log.Println("unauthorized access")
		return
	}

	buildingId, err := handler.UseCase.UpdateProgress(payloadService.DataDocument)
	if err != nil {
		log.Println("error update progress document, err: ", err)
		return
	}

	if payloadService.DataDocument.Status == "VERIFIED" {
		building := model.Building{
			ID:     buildingId,
			Status: "RESERVED",
		}

		payload := &model.PayloadServices{
			SecretToken:  os.Getenv("TOKEN_SERVICE"),
			DataBuilding: &building,
		}

		EmitEventToBuildingService(payload, handler.Redis, os.Getenv("UPDATE_BUILDING_CHANNEL"))
	}

}

func (handler *ApprovalHandlerImpl) GetDocument(ctx echo.Context) error {

	id := ctx.QueryParam("id")

	res, err := handler.UseCase.GetDocument(id)
	if err != nil {
		return ctx.JSON(http.StatusOK, model.Reponse{
			Status:  "failed",
			Message: "failed to fetch data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Reponse{
		Status:   "success",
		Message:  "success get data",
		Document: res,
	})

}
