package ports

import "pilotoAPI/internal/app/domain/model"

type ProductImpl interface {
	GetProduct(id int64) (*model.Product, error)
	GetProducts() ([]*model.Product, error)
	CreateProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(product *model.Product) (*model.Product, error)
	DeleteProduct(id int64) error
}
