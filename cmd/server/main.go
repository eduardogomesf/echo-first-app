package main

import (
	webserver "github.com/eduardogomesf/echo-first-app/internal/infra/http"
	"github.com/eduardogomesf/echo-first-app/internal/infra/http/handlers"
)

func main() {
	ws := webserver.NewWebServer()
	ws.AddHandler("/", "GET", handlers.Health)
	ws.AddHandler("/products", "GET", handlers.GetProducts)
	ws.AddHandler("/products", "POST", handlers.AddProduct)
	ws.AddHandler("/products/:name", "GET", handlers.GetProductByName)
	ws.Start(":3000")
}
