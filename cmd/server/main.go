package main

import (
	"log"

	configs "github.com/eduardogomesf/echo-first-app/cmd/config"
	webserver "github.com/eduardogomesf/echo-first-app/internal/infra/http"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/handlers"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/middlewares"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := configs.GetEnv("APP_PORT", "8080")

	ws := webserver.NewWebServer()

	// health
	healthHandler := handlers.HealthHandler{}
	ws.AddHandler("/", "GET", healthHandler.Health)

	// products
	productsHandler := handlers.ProductsHandler{}
	ws.AddHandler("/products", "POST", productsHandler.AddProduct, middlewares.UseAuthMiddleware())
	ws.AddHandler("/products", "GET", productsHandler.GetProducts)
	ws.AddHandler("/products/:name", "GET", productsHandler.GetProductByName)

	// auth
	authHandler := handlers.AuthHandler{}
	ws.AddHandler("/login", "POST", authHandler.Login)

	ws.Start(":" + port)
}
