package adapters

import (
	"pilotoAPI/internal/app/domain/command"
	"pilotoAPI/internal/app/domain/model"
)

type ProductService struct {
	productCmmd *command.ProductCmmd
}

func NewProductService(productCmmd *command.ProductCmmd) *ProductService {
	return &ProductService{
		productCmmd: productCmmd,
	}
}

func (s *ProductService) GetProduct(id int64) (*model.Product, error) {
	output, err := s.productCmmd.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ProductService) GetProducts() ([]*model.Product, error) {
	output, err := s.productCmmd.GetProducts()
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ProductService) CreateProduct(input *model.Product) (*model.Product, error) {
	output, err := s.productCmmd.CreateProduct(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ProductService) UpdateProduct(input *model.Product) (*model.Product, error) {
	output, err := s.productCmmd.UpdateProduct(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.productCmmd.DeleteProduct(id)
}
