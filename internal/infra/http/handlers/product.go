package handlers

import (
	"net/http"

	"github.com/eduardogomesf/echo-first-app/internal/domain/entities"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/dtos"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

var products []entities.Product = []entities.Product{}

type ProductsHandler struct {
	dbConn *pgx.Conn
}

func NewProductsHandler(dbConn *pgx.Conn) *ProductsHandler {
	return &ProductsHandler{
		dbConn: dbConn,
	}
}

func (pc *ProductsHandler) GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func (pc *ProductsHandler) AddProduct(c echo.Context) error {
	dto := new(dtos.AddProductDto)

	if err := c.Bind(dto); err != nil {
		return err
	}

	p := entities.NewProduct(
		dto.Name,
		dto.Price,
		dto.Categories,
		dto.IsDisabled,
	)

	products = append(products, *p)

	return c.JSON(http.StatusOK, p)
}

func (pc *ProductsHandler) GetProductByName(c echo.Context) error {
	name := c.Param("name")

	for _, p := range products {
		if p.Name == name {
			return c.JSON(http.StatusOK, p)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
}
