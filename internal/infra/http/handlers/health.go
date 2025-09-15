package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	dbCon *pgx.Conn
}

func NewHealthController(dbCon *pgx.Conn) *HealthHandler {
	return &HealthHandler{
		dbCon: dbCon,
	}
}

func (hc *HealthHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "healthy"})
}

func (hc *HealthHandler) HealthZ(c echo.Context) error {
	err := hc.dbCon.Ping(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"db": "failed to ping database"})
	}

	return c.JSON(http.StatusOK, map[string]string{"db": "healthy"})
}
