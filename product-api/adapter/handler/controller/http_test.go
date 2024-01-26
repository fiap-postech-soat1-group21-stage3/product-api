// package controller_test

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"strings"
// 	"testing"

// 	"github.com/cucumber/godog"
// 	"github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/handler/controller"
// 	"github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/model"
// 	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/entity"
// 	mocks "github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/port/mocks"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	productModelInput = &model.ProductRequestDTO{
// 		Name:        "Batata frita",
// 		Description: "Batata frita temperada",
// 		Category:    "sides",
// 		Price:       10.00,
// 	}

// 	productEntityInput = &entity.Product{
// 		Name:        "Batata frita",
// 		Description: "Batata frita temperada",
// 		Category:    "sides",
// 		Price:       10.00,
// 	}

// 	productEntityOutput = &entity.Product{
// 		ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 		Name:        "Batata frita",
// 		Description: "Batata frita temperada",
// 		Category:    "sides",
// 		Price:       10.00,
// 	}

// 	productModelOutput = &model.ProductResponseDTO{
// 		ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 		Name:        "Batata frita",
// 		Description: "Batata frita temperada",
// 		Category:    "sides",
// 		Price:       10.00,
// 	}

// 	path = "/product/8c2b51bf-7b4c-4a4b-a024-f283576cf190"

// 	patchEntityInput = &entity.Product{
// 		ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 		Description: "Batata frita com molho",
// 		Price:       15.00,
// 	}

// 	// patchModelInput = &model.ProductRequestDTO{
// 	// 	Description: "Batata frita com molho",
// 	// 	Price:       15.00,
// 	// }

// 	id = &entity.Product{
// 		ID: uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 	}

// 	products = &entity.ProductResponseList{
// 		Result: []*entity.Product{
// 			{
// 				ID:          uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 				Name:        "Batata frita",
// 				Description: "Batata frita temperada",
// 				Category:    "sides",
// 				Price:       10.00,
// 			},
// 		},
// 		Count: 1,
// 	}
// )

// type handlerContext struct {
// 	handler   *controller.Handler
// 	w         *httptest.ResponseRecorder
// 	req       *http.Request
// 	err       error
// 	body      []byte
// 	productID string
// }

// // var productModelInput *model.ProductRequestDTO

// func convertStringToProductCategory(str string) (model.ProductCategory, error) {
// 	switch str {
// 	case "burgers":
// 		return model.Burgers, nil
// 	case "sides":
// 		return model.Sides, nil
// 	case "beverage":
// 		return model.Beverage, nil
// 	case "sweets":
// 		return model.Sweets, nil
// 	default:
// 		return "", fmt.Errorf("Invalid ProductCategory value: %s", str)
// 	}
// }

// func (h *handlerContext) theFollowingProductDetails(table *godog.Table) error {
// 	price, err := strconv.ParseFloat(table.Rows[1].Cells[3].Value, 64)
// 	if err != nil {
// 		return err
// 	}

// 	category, err := convertStringToProductCategory(table.Rows[1].Cells[2].Value)
// 	if err != nil {
// 		return err
// 	}

// 	productModelInput = &model.ProductRequestDTO{
// 		Name:        table.Rows[1].Cells[0].Value,
// 		Description: table.Rows[1].Cells[1].Value,
// 		Category:    category,
// 		Price:       price,
// 	}
// 	return nil
// }

// func (h *handlerContext) aRequestIsMadeToCreateTheProduct() error {
// 	jsonBytes, err := json.Marshal(productModelInput)
// 	if err != nil {
// 		return err
// 	}

// 	h.req = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBytes))
// 	h.w = httptest.NewRecorder()

// 	ctxGin, _ := gin.CreateTestContext(h.w)
// 	ctxGin.Request = h.req

// 	usecaseMock := &mocks.ProductUseCase{}
// 	usecaseMock.On("Create", ctxGin, productEntityInput).Return(productEntityOutput, h.err).Once()

// 	h.handler = controller.NewHandler(usecaseMock)
// 	h.handler.Create(ctxGin)

// 	res := h.w.Result()
// 	defer res.Body.Close()
// 	h.body = h.w.Body.Bytes()

// 	return nil
// }

// func (h *handlerContext) anExistingProductWithID(id string) error {
// 	h.productID = id
// 	return nil
// }

// func (h *handlerContext) aRequestIsMadeToUpdateTheProduct() error {
// 	h.req = httptest.NewRequest(http.MethodPut, "/product/8c2b51bf-7b4c-4a4b-a024-f283576cf190", nil)
// 	h.w = httptest.NewRecorder()

// 	ctxGin, _ := gin.CreateTestContext(h.w)
// 	ctxGin.Request = h.req

// 	usecaseMock := &mocks.ProductUseCase{}
// 	usecaseMock.On("Update", ctxGin, patchEntityInput).Return(h.err).Once()

// 	h.handler = controller.NewHandler(usecaseMock)
// 	h.handler.Update(ctxGin)

// 	res := h.w.Result()
// 	defer res.Body.Close()
// 	h.body = h.w.Body.Bytes()

// 	return nil
// }

// func (h *handlerContext) aRequestIsMadeToDeleteTheProduct() error {
// 	h.req = httptest.NewRequest(http.MethodDelete, path, nil)
// 	h.w = httptest.NewRecorder()

// 	ctxGin, _ := gin.CreateTestContext(h.w)
// 	ctxGin.Request = h.req

// 	usecaseMock := &mocks.ProductUseCase{}
// 	usecaseMock.On("Delete", ctxGin, id).Return(h.err).Once()

// 	h.handler = controller.NewHandler(usecaseMock)
// 	h.handler.Delete(ctxGin)

// 	res := h.w.Result()
// 	defer res.Body.Close()
// 	h.body = h.w.Body.Bytes()

// 	return nil
// }

// func (h *handlerContext) aRequestIsMadeToGetTheListofProducts() error {
// 	h.req = httptest.NewRequest(http.MethodGet, "/", nil)
// 	h.w = httptest.NewRecorder()

// 	ctxGin, _ := gin.CreateTestContext(h.w)
// 	ctxGin.Request = h.req

// 	usecaseMock := &mocks.ProductUseCase{}
// 	usecaseMock.On("GetProducts", ctxGin).Return(products, h.err).Once()

// 	h.handler = controller.NewHandler(usecaseMock)
// 	h.handler.GetProducts(ctxGin)

// 	res := h.w.Result()
// 	defer res.Body.Close()
// 	h.body = h.w.Body.Bytes()

// 	return nil
// }

// func (h *handlerContext) theResponseShouldHaveStatusCode(statusCode int) error {
// 	return assertExpectedAndActual(assert.Equal, statusCode, h.w.Code, "status code")
// }

// func (h *handlerContext) theResponseBodyShouldMatchTheExpectedProductDetails() error {
// 	wantGot, err := json.Marshal(productModelOutput)
// 	if err != nil {
// 		return err
// 	}

// 	return assertExpectedAndActual(assert.Equal, string(wantGot), strings.TrimSuffix(h.w.Body.String(), "\n"), "response body")
// }

// // assertExpectedAndActual is a helper function to allow the step function to call
// // assertion functions where you want to compare an expected and an actual value.
// func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
// 	var t asserter
// 	a(&t, expected, actual, msgAndArgs...)
// 	return t.err
// }

// type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// // asserter is used to be able to retrieve the error reported by the called assertion
// type asserter struct {
// 	err error
// }

// // Errorf is used by the called assertion to report an error
// func (a *asserter) Errorf(format string, args ...interface{}) {
// 	a.err = fmt.Errorf(format, args...)
// }

// func TestFeatures(t *testing.T) {
// 	suite := godog.TestSuite{
// 		Name:                 "http",
// 		ScenarioInitializer:  InitializeScenario,
// 		TestSuiteInitializer: InitializeTestSuite,
// 		Options: &godog.Options{
// 			Format:   "pretty",
// 			Paths:    []string{"../../../features/http.feature"},
// 			TestingT: t,
// 		},
// 	}

// 	if suite.Run() != 0 {
// 		t.Fatal("non-zero status returned, failed to run feature tests")
// 	}
// }

// func InitializeScenario(s *godog.ScenarioContext) {
// 	h := &handlerContext{}
// 	s.Step(`^the following product details`, h.theFollowingProductDetails)
// 	s.Step(`^a request is made to create the product$`, h.aRequestIsMadeToCreateTheProduct)
// 	s.Step(`^the response should have status code (\d+)$`, h.theResponseShouldHaveStatusCode)
// 	s.Step(`^the response body should match the expected product details`, h.theResponseBodyShouldMatchTheExpectedProductDetails)
// 	s.Given(`^an existing product with ID "([^"]*)`, h.anExistingProductWithID)
// 	s.Step(`^a request is made to update the product$`, h.aRequestIsMadeToUpdateTheProduct)
// 	s.Step(`^a request is made to delete the product$`, h.aRequestIsMadeToDeleteTheProduct)
// 	s.Step(`^a request is made to get the list of products$`, h.aRequestIsMadeToGetTheListofProducts)

// 	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
// 		h.err = nil
// 		return ctx, nil
// 	})

// 	s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
// 		return ctx, nil
// 	})
// }

// func InitializeTestSuite(ctx *godog.TestSuiteContext) {}
