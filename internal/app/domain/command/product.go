package command

import (
	"pilotoAPI/internal/app/domain/model"
	"pilotoAPI/internal/app/domain/repository"
)

type ProductCmmd struct {
	repo repository.ProductRepository
}

func NewProductCommand(repo repository.ProductRepository) *ProductCmmd {
	return &ProductCmmd{repo}
}

func (s *ProductCmmd) GetProduct(id int64) (*model.Product, error) {
	product, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}, nil
}

func (s *ProductCmmd) GetProducts() ([]*model.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	productsModel := make([]*model.Product, len(products))
	for idx, product := range products {
		productsModel[idx] = &model.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
		}
	}

	return productsModel, nil
}

func (s *ProductCmmd) CreateProduct(product *model.Product) (*model.Product, error) {
	productDB := &repository.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}

	err := s.repo.Create(productDB)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          productDB.ID,
		Name:        productDB.Name,
		Description: productDB.Description,
		Price:       productDB.Price,
		Quantity:    productDB.Quantity,
	}, nil
}

func (s *ProductCmmd) UpdateProduct(product *model.Product) (*model.Product, error) {
	productDB := &repository.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}

	err := s.repo.Update(productDB)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          productDB.ID,
		Name:        productDB.Name,
		Description: productDB.Description,
		Price:       productDB.Price,
		Quantity:    productDB.Quantity,
	}, nil
}

func (s *ProductCmmd) DeleteProduct(id int64) error {
	return s.repo.Delete(id)
}
