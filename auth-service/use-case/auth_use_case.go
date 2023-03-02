package usecase

import "auth-service/models"

type AuthUseCase interface {
	Auth(payload *models.LoginPayload) (string, error)
}
