package main

import (
	"echo-first-app/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Health)
	e.GET("/products", handlers.GetProducts)
	e.POST("/products", handlers.AddProduct)
	e.GET("/products/:name", handlers.GetProductByName)
	e.Logger.Fatal(e.Start(":3000"))
}
