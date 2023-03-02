package usecase

import (
	"auth-service/models"
	"auth-service/utils"
	"errors"

	authError "auth-service/errors"
	"auth-service/repository"
)

type AuthUseCaseImpl struct {
	Repo repository.AuthRepo
}

func NewAuthUseCase(repo repository.AuthRepo) AuthUseCase {
	return &AuthUseCaseImpl{
		Repo: repo,
	}
}

func (useCase *AuthUseCaseImpl) Auth(payload *models.LoginPayload) (string, error) {

	res, err := useCase.Repo.Auth(payload)
	if err != nil {
		return "", err
	}

	if ok := utils.CheckPasswordHash(payload.Password, res.Password); !ok {
		return "", errors.New(authError.INCORRECT_USERNAME_OR_PASSWORD)
	}

	token, err := utils.GenerateJWT(res)
	if err != nil {
		return "", err
	}

	return token, nil

}
