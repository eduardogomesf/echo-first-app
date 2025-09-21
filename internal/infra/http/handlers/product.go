package handlers

import (
	"context"
	"fmt"
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

	insertQuery := "INSERT INTO products (id, name, price, categories, is_disabled, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := pc.dbConn.Exec(
		context.Background(),
		insertQuery,
		p.Id, p.Name, p.Price, p.Categories, p.IsDisabled, p.CreatedAt, p.UpdatedAt,
	)

	if err != nil {
		errMessage := fmt.Errorf("error saving data in the database %s", err)
		fmt.Println(errMessage)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Fail to create product"})
	}

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
