package command

import (
	"github.com/stretchr/testify/assert"
	"pilotoAPI/internal/app/domain/model"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "pilotoAPI/internal/app/domain/mocks"
	"pilotoAPI/internal/app/domain/repository"
)

func TestGetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	svc := NewProductCommand(repo)

	t.Run("GetProduct", func(t *testing.T) {
		repo.EXPECT().Get(gomock.Any()).Return(&repository.Product{
			ID:          1,
			Name:        "Test",
			Description: "Test",
			Price:       1,
			Quantity:    1,
		}, nil)

		_, err := svc.GetProduct(1)
		assert.NoError(t, err)
	})

	t.Run("GetProductError", func(t *testing.T) {
		repo.EXPECT().Get(gomock.Any()).Return(nil, assert.AnError)

		_, err := svc.GetProduct(1)
		assert.Error(t, err)
	})
}

func TestGetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	svc := NewProductCommand(repo)

	t.Run("GetProducts", func(t *testing.T) {
		repo.EXPECT().GetAll().Return([]*repository.Product{
			{
				ID:          1,
				Name:        "Test",
				Description: "Test",
				Price:       1,
				Quantity:    1,
			},
		}, nil)

		_, err := svc.GetProducts()
		assert.NoError(t, err)
	})

	t.Run("GetProductsError", func(t *testing.T) {
		repo.EXPECT().GetAll().Return(nil, assert.AnError)

		_, err := svc.GetProducts()
		assert.Error(t, err)
	})
}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	svc := NewProductCommand(repo)

	t.Run("CreateProduct", func(t *testing.T) {
		repo.EXPECT().Create(gomock.Any()).Return(nil)

		_, err := svc.CreateProduct(&model.Product{
			Name:        "Test",
			Description: "Test",
			Price:       1,
			Quantity:    1,
		})
		assert.NoError(t, err)
	})

	t.Run("CreateProductError", func(t *testing.T) {
		repo.EXPECT().Create(gomock.Any()).Return(assert.AnError)

		_, err := svc.CreateProduct(&model.Product{
			Name:        "Test",
			Description: "Test",
			Price:       1,
			Quantity:    1,
		})
		assert.Error(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	svc := NewProductCommand(repo)

	t.Run("UpdateProduct", func(t *testing.T) {
		repo.EXPECT().Update(gomock.Any()).Return(nil)

		_, err := svc.UpdateProduct(&model.Product{
			ID:          1,
			Name:        "Test",
			Description: "Test",
			Price:       1,
			Quantity:    1,
		})
		assert.NoError(t, err)
	})

	t.Run("UpdateProductError", func(t *testing.T) {
		repo.EXPECT().Update(gomock.Any()).Return(assert.AnError)

		_, err := svc.UpdateProduct(&model.Product{
			ID:          1,
			Name:        "Test",
			Description: "Test",
			Price:       1,
			Quantity:    1,
		})
		assert.Error(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	svc := NewProductCommand(repo)

	t.Run("DeleteProduct", func(t *testing.T) {
		repo.EXPECT().Delete(gomock.Any()).Return(nil)

		err := svc.DeleteProduct(1)
		assert.NoError(t, err)
	})

	t.Run("DeleteProductError", func(t *testing.T) {
		repo.EXPECT().Delete(gomock.Any()).Return(assert.AnError)

		err := svc.DeleteProduct(1)
		assert.Error(t, err)
	})
}
