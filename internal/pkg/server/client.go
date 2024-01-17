package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpServer interface {
	Serve(*echo.Echo, string) error
}

type HttpServerSvc struct {
	server *http.Server
}

func NewHTTPServer() *HttpServerSvc {
	return &HttpServerSvc{}
}
