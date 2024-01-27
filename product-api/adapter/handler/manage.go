package manage

import (
	p "github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type Manage struct {
	product apps
}

type UseCases struct {
	Product port.ProductUseCase
}

func New(uc *UseCases) *Manage {

	productHandler := p.NewHandler(uc.Product)

	return &Manage{
		product: productHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.product.RegisterRoutes(group)
}
