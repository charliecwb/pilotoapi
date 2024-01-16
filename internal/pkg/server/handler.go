package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HandlerProvider func(context context.Context, router fiber.Router)

func (s *HttpServerSvc) NewHandler() http.HandlerFunc {
	return s.newHandler("")
}

func (s *HttpServerSvc) newHandler(healthPattern string) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(rw, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		if healthPattern != "" {
			if r.URL.Path != healthPattern {
				http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
		}

		rw.Header().Set("Content-Type", "application/json")
	}
}
