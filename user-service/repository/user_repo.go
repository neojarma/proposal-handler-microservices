package repository

import "user-service/models"

type UserRepo interface {
	Register(request *models.User) error
}
