package repository

import (
	authError "auth-service/errors"
	model "auth-service/models"
	"errors"

	"gorm.io/gorm"
)

type AuthRepoImpl struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &AuthRepoImpl{
		DB: db,
	}
}

func (repo *AuthRepoImpl) Auth(request *model.LoginPayload) (*model.User, error) {
	model := new(model.User)

	res := repo.DB.Where("username = ?", request.Username).First(model)
	if res.RowsAffected == 0 {
		return nil, errors.New(authError.INCORRECT_USERNAME_OR_PASSWORD)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return model, nil
}
