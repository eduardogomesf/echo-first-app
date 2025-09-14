package middlewares

import (
	"log"

	configs "github.com/eduardogomesf/echo-first-app/cmd/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UseAuthMiddleware() echo.MiddlewareFunc {
	secret := configs.GetEnv("JWT_SECRET", "")

	if secret == "" {
		log.Fatal("JWT Secret not found in ENV")
	}

	return echojwt.JWT([]byte(secret))
}
