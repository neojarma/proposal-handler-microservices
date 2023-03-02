package middleware

import (
	"building-service/models"
	"net/http"

	"building-service/utils"

	"github.com/labstack/echo/v4"
)

type AuthMiddlewareImpl struct{}

func NewAuthMiddleware() AuthMiddleware {
	return &AuthMiddlewareImpl{}
}

func (middleware *AuthMiddlewareImpl) IsLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("access_token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Reponse{
				Status:  "failed",
				Message: "missing access token",
			})
		}

		token := cookie.Value

		_, ok := utils.VerifyJWT(token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, models.Reponse{
				Status:  "failed",
				Message: "unauthorized access",
			})
		}

		return next(c)

	}
}

func (middleware *AuthMiddlewareImpl) IsStaff(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("access_token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Reponse{
				Status:  "failed",
				Message: "missing access token",
			})
		}

		token := cookie.Value

		role, ok := utils.VerifyJWT(token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, models.Reponse{
				Status:  "failed",
				Message: "unauthorized access",
			})
		}

		if role != "STAFF" {
			return c.JSON(http.StatusUnauthorized, models.Reponse{
				Status:  "failed",
				Message: "unauthorized access",
			})
		}

		return next(c)

	}
}
