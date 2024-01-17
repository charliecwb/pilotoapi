package routes

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
	"net/http"
	"pilotoAPI/internal/app/domain/model"
	"pilotoAPI/internal/pkg/errors"
	"strconv"
)

const (
	errFalhaConsultarDados = errors.ErrFalhaConsultarDados
	errParameterInvalido   = errors.ErrParameterInvalido
	errRequestBodyInvalido = errors.ErrRequestBodyInvalido
	errFalhaEnviarDados    = errors.ErrFalhaEnviarDados
)

type FetchProductByIdRequest struct {
	ID int64 `query:"id" params:"id" validate:"nonzero" json:"id"`
}

type ProductResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	CreatedAt   string  `json:"createdAt"`
}

func (h *Handler) fetchProduct(ctx echo.Context) error {
	req := new(FetchProductByIdRequest)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errParameterInvalido})
	}

	req.ID = int64(id)

	if err := validator.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errParameterInvalido})
	}

	resp, err := h.productSvc.GetProduct(req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse{err.Error(), errFalhaConsultarDados})
	}

	return ctx.JSON(http.StatusOK, parseProductToResponse(resp))
}

func (h *Handler) fetchAllProducts(ctx echo.Context) error {
	resp, err := h.productSvc.GetProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse{err.Error(), errFalhaConsultarDados})
	}

	return ctx.JSON(http.StatusOK, parseProductToListResponse(resp))
}

func (h *Handler) createProduct(ctx echo.Context) error {
	product := new(model.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errRequestBodyInvalido})
	}

	if err := validator.Validate(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errParameterInvalido})
	}

	resp, err := h.productSvc.CreateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse{err.Error(), errFalhaEnviarDados})
	}

	return ctx.JSON(http.StatusCreated, parseProductToResponse(resp))
}

func (h *Handler) updateProduct(ctx echo.Context) error {
	product := new(model.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errRequestBodyInvalido})
	}

	if err := validator.Validate(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errParameterInvalido})
	}

	resp, err := h.productSvc.UpdateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse{err.Error(), errFalhaEnviarDados})
	}

	return ctx.JSON(http.StatusOK, parseProductToResponse(resp))
}

func (h *Handler) deleteProduct(ctx echo.Context) error {
	req := new(FetchProductByIdRequest)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errParameterInvalido})
	}

	req.ID = int64(id)

	if err := validator.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ErrorResponse{err.Error(), errParameterInvalido})
	}

	err = h.productSvc.DeleteProduct(req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.ErrorResponse{err.Error(), errFalhaEnviarDados})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func parseProductToListResponse(products []*model.Product) []*ProductResponse {
	productsResponse := make([]*ProductResponse, len(products))
	for idx, product := range products {
		productsResponse[idx] = parseProductToResponse(product)
	}

	return productsResponse
}

func parseProductToResponse(product *model.Product) *ProductResponse {
	return &ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
