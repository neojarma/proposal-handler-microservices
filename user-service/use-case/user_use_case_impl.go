package usecase

import (
	"user-service/models"
	repo "user-service/repository"

	utils "user-service/utils"
)

type UserUseCaseImpl struct {
	Repo repo.UserRepo
}

func NewUserUseCase(repo repo.UserRepo) UserUseCase {
	return &UserUseCaseImpl{
		Repo: repo,
	}
}

func (useCase *UserUseCaseImpl) Regist(request *models.User) error {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = hashedPass
	err = useCase.Repo.Register(request)
	if err != nil {
		return err
	}

	return nil
}
