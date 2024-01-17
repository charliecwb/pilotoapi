package server

import (
	"github.com/labstack/echo/v4"
)

type HttpServer interface {
	Serve(ec *echo.Echo, url string) error
}

type HttpServerSvc struct {
}

func NewHTTPServer() *HttpServerSvc {
	return &HttpServerSvc{}
}
