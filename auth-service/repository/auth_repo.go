package repository

import model "auth-service/models"

type AuthRepo interface {
	Auth(request *model.LoginPayload) (*model.User, error)
}
