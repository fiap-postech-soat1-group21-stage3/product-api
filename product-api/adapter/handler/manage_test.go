// package manage

// import (
// 	"testing"

// 	"github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/handler/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestManageRegisterRoutes(t *testing.T) {
// 	appMock := mocks.NewApps(t)

// 	t.Run("when everything goes ok; should call apps", func(t *testing.T) {
// 		appMock.On("RegisterRoutes", mock.AnythingOfType("*gin.RouterGroup")).Return().Once()
// 		m := &Manage{
// 			product: appMock,
// 		}
// 		m.RegisterRoutes(nil)

// 		appMock.AssertExpectations(t)
// 	})
// }

// func TestNew(t *testing.T) {
// 	got := New(&UseCases{})
// 	assert.NotNil(t, got.product)
// }
