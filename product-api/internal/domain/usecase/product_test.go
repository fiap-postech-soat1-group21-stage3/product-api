package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/entity"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/port/mocks"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	defaultID  = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")
	ctxDefault = context.Background()

	given = &entity.Product{
		Name:        "Burguer de planta",
		Description: "Burguer com carne vegetal",
		Category:    entity.Burguers,
		Price:       10,
	}

	wanted = &entity.Product{
		ID:          defaultID,
		Name:        "Burguer de planta",
		Description: "Burguer com carne vegetal",
		Category:    entity.Burguers,
		Price:       10,
	}

	update = &entity.Product{
		ID:          defaultID,
		Name:        "Burguer de planta",
		Description: "Burguer com carne vegetal",
		Category:    entity.Burguers,
		Price:       10,
	}

	delete = &entity.Product{
		ID: defaultID,
	}

	productList = &entity.ProductResponseList{
		Result: []*entity.Product{
			{
				ID:          defaultID,
				Name:        "Burguer de planta",
				Description: "Burguer com carne vegetal",
				Category:    entity.Burguers,
				Price:       10,
			},
			{
				ID:          uuid.MustParse("b4dacf92-7000-4523-9fab-166212aaa92d"),
				Name:        "Burguer de tofu",
				Description: "Burguer com tofu",
				Category:    entity.Burguers,
				Price:       10,
			},
		},
		Count: 2,
	}
)

func Test_CreateProduct(t *testing.T) {
	t.Run("when everything goes as expected; should return a product and no error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		repo.On("Create", ctxDefault, given).Return(wanted, nil).Once()

		got, err := service.Create(ctxDefault, given)
		assert.NoError(t, err)
		assert.Equal(t, wanted, got)
		repo.AssertExpectations(t)
	})

	t.Run("when repo returns error; should propagate the error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		wantError := errors.New("error")

		repo.On("Create", ctxDefault, given).Return(nil, wantError).Once()

		got, err := service.Create(ctxDefault, given)
		assert.ErrorIs(t, err, wantError)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}

func Test_UpdateProduct(t *testing.T) {
	t.Run("when everything goes as expected; should return no error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		repo.On("Update", ctxDefault, update).Return(nil).Once()

		err := service.Update(ctxDefault, update)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("when repo returns error; should propagate the error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		wantError := errors.New("error")

		repo.On("Update", ctxDefault, update).Return(wantError).Once()

		err := service.Update(ctxDefault, update)
		assert.ErrorIs(t, err, wantError)
		repo.AssertExpectations(t)
	})
}

func Test_DeleteProduct(t *testing.T) {
	t.Run("when everything goes as expected; should return no error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		repo.On("Delete", ctxDefault, delete).Return(nil).Once()

		err := service.Delete(ctxDefault, delete)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("when repo returns error; should propagate the error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		wantError := errors.New("error")

		repo.On("Delete", ctxDefault, delete).Return(wantError).Once()

		err := service.Delete(ctxDefault, delete)
		assert.ErrorIs(t, err, wantError)
		repo.AssertExpectations(t)
	})
}

func TestRetrieveListProducts(t *testing.T) {
	t.Run("when everything goes as expected; should return a product list and no error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		repo.On("GetProducts", ctxDefault).Return(productList, nil).Once()

		got, err := service.GetProducts(ctxDefault)
		assert.NoError(t, err)
		assert.Equal(t, got, productList)
		repo.AssertExpectations(t)
	})
	t.Run("when repository returns error; should send forward the error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		service := usecase.NewProductUseCase(repo)

		wantedErr := errors.New("error")
		repo.On("GetProducts", ctxDefault).Return(nil, wantedErr).Once()

		got, err := service.GetProducts(ctxDefault)
		assert.ErrorIs(t, err, wantedErr)
		assert.Nil(t, got)
		repo.AssertExpectations(t)
	})
}
