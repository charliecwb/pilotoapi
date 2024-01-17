package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"pilotoAPI/internal/pkg/errors"
)

type HandlerProvider func(context context.Context, router fiber.Router)

type ErrorResponse errors.ErrorResponse
