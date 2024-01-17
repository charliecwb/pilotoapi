package ports

import "pilotoAPI/internal/app/domain/model"

//go:generate mockgen -source=products.go -destination=./../mocks/product_port.go -package=ports
type ProductImpl interface {
	GetProduct(id int64) (*model.Product, error)
	GetProducts() ([]*model.Product, error)
	CreateProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(product *model.Product) (*model.Product, error)
	DeleteProduct(id int64) error
}
