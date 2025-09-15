package handlers

import (
	"fmt"
	"net/http"
	"time"

	configs "github.com/eduardogomesf/echo-first-app/cmd/config"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/dtos"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (ac *AuthHandler) Login(c echo.Context) error {
	dto := new(dtos.LoginDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dto.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := configs.GetEnv("JWT_SECRET", "")

	if secret == "" {
		fmt.Println("No secret found during token generation. Login has failed")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal Server Error",
		})
	}

	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return err
	}

	response := dtos.LoginResponseDto{
		Token: tokenStr,
	}

	return c.JSON(http.StatusOK, response)
}
