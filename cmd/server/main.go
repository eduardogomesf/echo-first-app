package main

import (
	webserver "github.com/eduardogomesf/echo-first-app/internal/infra/http"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/handlers"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/middlewares"
)

func main() {
	ws := webserver.NewWebServer()

	ws.AddHandler("/", "GET", handlers.Health)

	ws.AddHandler("/products", "POST", handlers.AddProduct, middlewares.UseAuthMiddleware())
	ws.AddHandler("/products", "GET", handlers.GetProducts)
	ws.AddHandler("/products/:name", "GET", handlers.GetProductByName)

	ws.AddHandler("/login", "POST", handlers.Login)

	ws.Start(":3000")
}
