package handlers

import (
	"net/http"

	"github.com/eduardogomesf/echo-first-app/internal/domain/entities"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/dtos"
	"github.com/labstack/echo/v4"
)

var products []entities.Product = []entities.Product{}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func AddProduct(c echo.Context) error {
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

func GetProductByName(c echo.Context) error {
	name := c.Param("name")

	for _, p := range products {
		if p.Name == name {
			return c.JSON(http.StatusOK, p)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
}
