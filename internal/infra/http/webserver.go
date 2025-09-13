package webserver

import (
	"errors"
	"fmt"
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	Route       string
	Method      string
	HandlerFunc echo.HandlerFunc
}

type WebServer struct {
	server   *echo.Echo
	handlers []Handler
}

func NewWebServer() *WebServer {
	return &WebServer{
		server:   echo.New(),
		handlers: []Handler{},
	}
}

func (ws *WebServer) AddHandler(route string, method string, handler echo.HandlerFunc) error {
	allowedMethods := []string{"POST", "GET"} // TO DO: add more methods

	if !slices.Contains(allowedMethods, method) {
		return fmt.Errorf("method %s not allowed", method)
	}

	if route == "" {
		return errors.New("route can not be empty")
	}

	if handler == nil {
		return errors.New("handler must be a handler function")
	}

	h := new(Handler)

	h.HandlerFunc = handler
	h.Route = route
	h.Method = method

	ws.handlers = append(ws.handlers, *h)

	switch method {
	case "POST":
		ws.server.POST(route, handler)
	case "GET":
		ws.server.GET(route, handler)
	}

	return nil
}

func (ws *WebServer) applyGlobalMiddlewares() {
	ws.server.Use(middleware.Logger())
	ws.server.Use(middleware.Recover()) // TO DO: perform redundancy tests with this middleware disabled
}

func (ws *WebServer) Start(port string) {
	ws.applyGlobalMiddlewares()

	err := ws.server.Start(port)

	if err != nil {
		ws.server.Logger.Fatal(err)
	}

	fmt.Printf("Server running on port %s", port)
}
