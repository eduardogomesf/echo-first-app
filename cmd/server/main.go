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

	ws.AddHandler("/", "GET", handlers.Health)

	ws.AddHandler("/products", "POST", handlers.AddProduct, middlewares.UseAuthMiddleware())
	ws.AddHandler("/products", "GET", handlers.GetProducts)
	ws.AddHandler("/products/:name", "GET", handlers.GetProductByName)

	ws.AddHandler("/login", "POST", handlers.Login)

	ws.Start(":" + port)
}
