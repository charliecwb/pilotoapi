package routes

import (
	"log"
	"pilotoAPI/internal/app/domain/ports"
)

type Handler struct {
	productSvc ports.ProductImpl
	log        *log.Logger
}

func NewHandler(
	productSvc ports.ProductImpl,
	log *log.Logger,
) *Handler {
	return &Handler{
		productSvc: productSvc,
		log:        log,
	}
}
