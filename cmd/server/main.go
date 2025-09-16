package main

import (
	"context"
	"fmt"
	"log"

	configs "github.com/eduardogomesf/echo-first-app/cmd/config"
	webserver "github.com/eduardogomesf/echo-first-app/internal/infra/http"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/handlers"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/middlewares"
	pgx "github.com/jackc/pgx/v5"
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
	dbConnStr := configs.GetEnv("DB_URL", "postgres://postgres:password@localhost:5432/catalog")

	ctx := context.Background()

	ws := webserver.NewWebServer()

	dbConn, err := pgx.Connect(ctx, dbConnStr)

	if err != nil {
		log.Fatal(fmt.Errorf("failed to open DB: %s", err))
	}

	err = dbConn.Ping(ctx)

	if err != nil {
		log.Fatal(fmt.Errorf("failed to ping DB: %s", err))
	}

	defer dbConn.Close(ctx)

	// health
	healthHandler := handlers.NewHealthController(dbConn)
	ws.AddHandler("/health", "GET", healthHandler.Health)
	ws.AddHandler("/healthz", "GET", healthHandler.HealthZ)

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
