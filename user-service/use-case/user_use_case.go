package usecase

import "user-service/models"

type UserUseCase interface {
	Regist(request *models.User) error
}
