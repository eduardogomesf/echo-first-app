package middlewares

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UseAuthMiddleware() echo.MiddlewareFunc {
	return echojwt.JWT([]byte("any-secret")) // TO DO: Retrieve secret from env
}
