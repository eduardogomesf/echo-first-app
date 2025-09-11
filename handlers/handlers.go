package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var products []Product = []Product{}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "server running"})
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func AddProduct(c echo.Context) error {
	p := new(Product)

	if err := c.Bind(p); err != nil {
		return err
	}

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
