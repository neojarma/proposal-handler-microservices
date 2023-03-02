package handler

import (
	"encoding/json"
	"log"
	"user-service/models"
	usecase "user-service/use-case"
	"user-service/utils"
)

type UserHandlerImpl struct {
	UseCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) UserHandler {
	return &UserHandlerImpl{
		UseCase: useCase,
	}
}

func (handler *UserHandlerImpl) Regist(payload string) {
	payloadServices := new(models.PayloadServices)
	err := json.Unmarshal([]byte(payload), payloadServices)
	if err != nil {
		log.Println("error while unmarshal payload, err:", err)
		return
	}

	if ok := utils.CheckTokenService(payloadServices.SecretToken); !ok {
		log.Println("unauthorized access")
		return
	}

	err = handler.UseCase.Regist(payloadServices.DataRegist)
	if err != nil {
		log.Println("error regist user, err:", err)
		return
	}
}
