package repository

import (
	"user-service/models"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}

func (repo *UserRepoImpl) Register(request *models.User) error {

	res := repo.DB.Omit("id", "role").Create(request)
	if res.Error != nil {
		return res.Error
	}

	return nil

}
