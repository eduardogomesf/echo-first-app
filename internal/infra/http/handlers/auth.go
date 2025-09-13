package handlers

import (
	"net/http"
	"time"

	"github.com/eduardogomesf/echo-first-app/internal/infra/http/dtos"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	dto := new(dtos.LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dto.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte("any-secret")) // TO DO: use env

	if err != nil {
		return err
	}

	response := dtos.LoginResponseDto{
		Token: tokenStr,
	}

	return c.JSON(http.StatusOK, response)
}
