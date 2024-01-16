package server

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HttpServer interface {
	Serve(*fiber.App, uint16) error
}

type HttpServerSvc struct {
	server *http.Server
}

func NewHTTPServer() *HttpServerSvc {
	return &HttpServerSvc{}
}
