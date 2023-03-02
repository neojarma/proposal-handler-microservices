package handler

import (
	"auth-service/models"
	useCase "auth-service/use-case"
	"auth-service/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandlerImpl struct {
	UseCase useCase.AuthUseCase
}

func NewAuthHandler(useCase useCase.AuthUseCase) AuthHandler {
	return &AuthHandlerImpl{
		UseCase: useCase,
	}
}

func (handler *AuthHandlerImpl) Auth(ctx echo.Context) error {

	model := new(models.LoginPayload)
	if err := ctx.Bind(model); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	token, err := handler.UseCase.Auth(model)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Reponse{
			Status:  "failed",
			Message: "failed to authenticate",
		})
	}

	cookie := utils.GenerateCookie("access_token", token)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, models.Reponse{
		Status:  "success",
		Message: "success to log in",
	})
}
