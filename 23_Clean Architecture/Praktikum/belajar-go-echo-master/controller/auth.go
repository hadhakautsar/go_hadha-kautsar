package controller

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secret := []byte("my-secret-key")

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			return secret, nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(ctx)
	}
}
