package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/adapter/model"
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/entity"
	mocks "github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createModelInput = &model.ProductRequestDTO{
		Name:        "Hamburguer",
		Description: "hamburguer de feijão",
		Category:    "hamburguer",
		Price:       20,
	}

	createEntityInput = &entity.Product{
		Name:        "Hamburguer",
		Description: "hamburguer de feijão",
		Category:    "hamburguer",
		Price:       20,
	}

	createModelOutput = &model.ProductResponseDTO{
		ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
		Name:        "Hamburguer",
		Description: "hamburguer de feijão",
		Category:    "hamburguer",
		Price:       20,
	}

	createEntityOutput = &entity.Product{
		ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
		Name:        "Hamburguer",
		Description: "hamburguer de feijão",
		Category:    "hamburguer",
		Price:       20,
	}

	retrieveEntityOutput = &entity.ProductResponseList{
		Result: []*entity.Product{
			{
				ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
				Name:        "Hamburguer",
				Description: "hamburguer de feijão",
				Category:    "hamburguer",
				Price:       20,
			},
		},
		Count: 1,
	}

	retrieveModelOutput = &model.ProductResponseList{
		Result: []*model.ProductResponseDTO{
			{
				ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf192"),
				Name:        "Hamburguer",
				Description: "hamburguer de feijão",
				Category:    "hamburguer",
				Price:       20,
			},
		},
		Count: 1,
	}
)

func TestCreate(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
		jsonBytes, err := json.Marshal(createModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		ctxGin, _ := gin.CreateTestContext(w)
		ctxGin.Request = req

		usecaseMock := mocks.NewProductUseCase(t)
		usecaseMock.On("Create", mock.AnythingOfType("*gin.Context"), createEntityInput).Return(createEntityOutput, nil).Once()

		handler := controller.NewHandler(usecaseMock)

		handler.Create(ctxGin)

		res := w.Result()
		defer res.Body.Close()
		got, err := json.Marshal(createModelOutput)

		assert.NoError(t, err)
		assert.EqualValues(t, strings.TrimSuffix(w.Body.String(), "\n"), string(got))
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when body is invalid; should return response 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{>}`)))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewProductUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		engine.POST("/product/", handler.Create)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestRetrieve(t *testing.T) {
	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		ctxGin, _ := gin.CreateTestContext(w)
		ctxGin.Request = req

		usecaseMock := mocks.NewProductUseCase(t)
		usecaseMock.On("GetProducts", ctxGin).Return(retrieveEntityOutput, nil).Once()

		handler := controller.NewHandler(usecaseMock)
		handler.GetProducts(ctxGin)

		res := w.Result()
		defer res.Body.Close()
		wantGot, err := json.Marshal(retrieveModelOutput)
		assert.NoError(t, err)

		assert.EqualValues(t, strings.TrimSuffix(w.Body.String(), "\n"), string(wantGot))
		assert.Equal(t, http.StatusOK, res.StatusCode)
		usecaseMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("when everything is ok, should return response 204", func(t *testing.T) {
		jsonBytes, err := json.Marshal(patchModelInput)
		if err != nil {
			return
		}
		req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonBytes))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewProductUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		usecaseMock.On("Update", mock.AnythingOfType("*gin.Context"), patchEntityInput).Return(nil).Once()

		assert.NoError(t, err)
		engine.PUT("/product/:id", handler.Update)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		usecaseMock.AssertExpectations(t)
	})

	t.Run("when body is invalid; should return response 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer([]byte(`{>}`)))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewProductUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		engine.PUT("/product/:id", handler.Update)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("when everything is ok, should return response 200", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, path, nil)
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewProductUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		usecaseMock.On("Delete", mock.AnythingOfType("*gin.Context"), deleteEntityInput).Return(nil).Once()

		engine.DELETE("/product/:id", handler.Delete)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		usecaseMock.AssertExpectations(t)
	})
}
