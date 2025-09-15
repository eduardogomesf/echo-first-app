package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func (hc *HealthHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "server running"})
}
