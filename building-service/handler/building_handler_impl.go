package handler

import (
	"building-service/models"
	usecase "building-service/use-case"
	"building-service/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BuildingHandlerImpl struct {
	UseCase usecase.BuildingUseCase
}

func NewBuildingHandler(useCase usecase.BuildingUseCase) BuildingHandler {
	return &BuildingHandlerImpl{
		UseCase: useCase,
	}
}

func (handler *BuildingHandlerImpl) GetBuildings(ctx echo.Context) error {
	var result []*models.Building
	var err error

	status := ctx.QueryParam("status")
	if status == "" {
		result, err = handler.UseCase.GetAll()
	} else {
		result, err = handler.UseCase.FilterByStatus(status)
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Reponse{
			Status:  "failed",
			Message: "failed to get data",
		})
	}

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:    "success",
		Message:   "success get data",
		Buildings: result,
	})
}

func (handler *BuildingHandlerImpl) NewBuilding(payload string) {
	payloadServices := new(models.PayloadServices)
	err := json.Unmarshal([]byte(payload), payloadServices)
	if err != nil {
		log.Println("error while unmarshal payload, err:", err)
		return
	}

	if ok := utils.CheckTokenService(payloadServices.SecretToken); !ok {
		log.Println("unathorized access")
		return
	}

	err = handler.UseCase.NewBuilding(payloadServices.DataBuilding)
	if err != nil {
		log.Println("error while creating new building, err:", err)
	}
}

func (handler *BuildingHandlerImpl) UpdateBuildingStatus(payload string) {
	payloadServices := new(models.PayloadServices)
	err := json.Unmarshal([]byte(payload), payloadServices)
	if err != nil {
		log.Println("error while unmarshal payload, err:", err)
		return
	}

	if ok := utils.CheckTokenService(payloadServices.SecretToken); !ok {
		log.Println("unathorized access")
		return
	}

	err = handler.UseCase.UpdateBuildingStatus(payloadServices.DataBuilding)
	if err != nil {
		log.Println("error while updating building status, err:", err)
	}
}
