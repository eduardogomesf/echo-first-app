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

type ProductsHandler struct {
	dbConn *pgx.Conn
}

func NewProductsHandler(dbConn *pgx.Conn) *ProductsHandler {
	return &ProductsHandler{
		dbConn: dbConn,
	}
}

func (pc *ProductsHandler) GetProducts(c echo.Context) error {
	rows, err := pc.dbConn.Query(context.Background(), "SELECT * FROM products")

	if err != nil {
		fmt.Println(fmt.Errorf("error retrieving products from database %s", err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Fail to retrieve products"})
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities.Product])

	if err != nil {
		fmt.Println(fmt.Errorf("error scanning products from query result to struct %s", err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Fail to retrieve products"})
	}

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

	// TO DO: do I need to validate if name has a valid value?

	row := pc.dbConn.QueryRow(context.Background(), "SELECT * FROM products where name = $1", name)

	product := entities.Product{}

	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Categories, &product.IsDisabled, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		// TO DO: differentiate between no match and actual error
		fmt.Println(fmt.Errorf("product not found by name %s: %s", name, err))
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
	}

	return c.JSON(http.StatusOK, product)
}
