package controller

import (
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase port.ProductUseCase
}

func NewHandler(u port.ProductUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	productRoute := routes.Group("/product")
	productRoute.POST("/", h.Create)
	productRoute.GET("/", h.GetProducts)
	productRoute.PUT("/:id", h.Update)
	productRoute.DELETE("/:id", h.Delete)
}
