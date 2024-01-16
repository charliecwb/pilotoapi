package repository

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var _ ProductRepository = new(ProductRepositoryImpl)

type Product struct {
	ID          int64          `gorm:"column:id"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	Price       float64        `gorm:"column:price"`
	Quantity    int64          `gorm:"column:quantity"`
	CreatedAt   mysql.NullTime `gorm:"column:created_at"`
}

func (t *Product) TableName() string {
	return "product"
}

type ProductRepository interface {
	Begin() *gorm.DB
	Commit() error
	Rollback()
	Create(*Product) error
	Update(*Product) error
	Get(int64) (*Product, error)
	Delete(int64) error
	GetAll() ([]*Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *ProductRepositoryImpl) Commit() error {
	return r.db.Commit().Error
}

func (r *ProductRepositoryImpl) Rollback() {
	r.db.Rollback()
}

func (r *ProductRepositoryImpl) Create(p *Product) error {
	return r.db.Create(p).Error
}

func (r *ProductRepositoryImpl) Update(p *Product) error {
	return r.db.Save(p).Error
}

func (r *ProductRepositoryImpl) Get(id int64) (*Product, error) {
	var p Product
	if err := r.db.Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepositoryImpl) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&Product{}).Error
}

func (r *ProductRepositoryImpl) GetAll() ([]*Product, error) {
	var products []*Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
