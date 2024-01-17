package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpServer interface {
	Serve(ec *echo.Echo, url string) error
}

type HttpServerSvc struct {
	server *http.Server
}

func NewHTTPServer() *HttpServerSvc {
	return &HttpServerSvc{}
}
