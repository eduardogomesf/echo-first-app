package main

import (
	webserver "github.com/eduardogomesf/echo-first-app/internal/infra/http"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/handlers"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/middlewares"
)

func main() {
	ws := webserver.NewWebServer()
	ws.AddHandler("/products", "POST", handlers.AddProduct, middlewares.UseAuthMiddleware())
	ws.AddHandler("/", "GET", handlers.Health)
	ws.AddHandler("/products", "GET", handlers.GetProducts)
	ws.AddHandler("/products/:name", "GET", handlers.GetProductByName)
	ws.Start(":3000")
}
