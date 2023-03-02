package handler

import (
	"api-gateway/models"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type APIHandlerImpl struct {
	RedisConn *redis.Client
}

func NewAPIHandler(conn *redis.Client) APIHandler {
	return &APIHandlerImpl{
		RedisConn: conn,
	}
}

func (handler *APIHandlerImpl) Register(ctx echo.Context) error {

	req := new(models.RegistPayload)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	channel := os.Getenv("USER_CHANNEL")
	payload := &models.PayloadServices{
		SecretToken: os.Getenv("TOKEN_SERVICE"),
		DataRegist:  req,
	}

	err := handler.emitEventToService(ctx.Request().Context(), payload, channel)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "failed send message to broker",
		})
	}

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:  "success",
		Message: "success send request",
	})
}

func (handler *APIHandlerImpl) ProposeDocument(ctx echo.Context) error {

	req := new(models.Document)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	newDocumentChannel := os.Getenv("NEW_DOCUMENT_CHANNEL")
	payload := &models.PayloadServices{
		SecretToken:  os.Getenv("TOKEN_SERVICE"),
		DataDocument: req,
	}

	err := handler.emitEventToService(ctx.Request().Context(), payload, newDocumentChannel)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "failed send message to broker",
		})
	}

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:  "success",
		Message: "success send request",
	})

}

func (handler *APIHandlerImpl) UpdateDocument(ctx echo.Context) error {

	req := new(models.Document)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	updateDocumentChannel := os.Getenv("UPDATE_DOCUMENT_CHANNEL")
	payload := &models.PayloadServices{
		SecretToken:  os.Getenv("TOKEN_SERVICE"),
		DataDocument: req,
	}

	err := handler.emitEventToService(ctx.Request().Context(), payload, updateDocumentChannel)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "failed send message to broker",
		})
	}

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:  "success",
		Message: "success send request",
	})

}

func (handler *APIHandlerImpl) NewBuilding(ctx echo.Context) error {
	req := new(models.Building)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	newBuildingChannel := os.Getenv("NEW_BUILDING_CHANNEL")
	payload := &models.PayloadServices{
		SecretToken:  os.Getenv("TOKEN_SERVICE"),
		DataBuilding: req,
	}

	err := handler.emitEventToService(ctx.Request().Context(), payload, newBuildingChannel)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "failed send message to broker",
		})
	}

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:  "success",
		Message: "success send request",
	})

}

func (handler *APIHandlerImpl) UpdateBuildingStatus(ctx echo.Context) error {
	req := new(models.Building)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	updateBuildingChannel := os.Getenv("UPDATE_BUILDING_CHANNEL")
	payload := &models.PayloadServices{
		SecretToken:  os.Getenv("TOKEN_SERVICE"),
		DataBuilding: req,
	}

	err := handler.emitEventToService(ctx.Request().Context(), payload, updateBuildingChannel)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "failed send message to broker",
		})
	}

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:  "success",
		Message: "success send request",
	})
}
