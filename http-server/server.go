package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"fmt"
)

type WebServer struct {
	Products string
}

func NewWebServer () *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.Start(":8080")
}

func main() {
	webserver := NewWebServer()
	webserver.Serve()
}


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Golang API RESPONSE </h1>")
}