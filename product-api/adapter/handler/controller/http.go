package controller

import (
	"net/http"

	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/adapter/model"
	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Create(ctx *gin.Context) {
	var input *model.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	domain := &entity.Product{
		Name:        input.Name,
		Description: input.Description,
		Category:    entity.Category(input.Category),
		Price:       input.Price,
	}

	res, err := h.useCase.Create(ctx, domain)
	if err != nil {
		return
	}

	output := &model.ProductResponseDTO{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Category:    model.ProductCategory(res.Category),
		Price:       res.Price,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, output)
}

func (h *Handler) Update(ctx *gin.Context) {
	var input *model.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}
	idParam := ctx.Param("id")

	id, _ := uuid.Parse(idParam)

	domain := &entity.Product{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Category:    entity.Category(input.Category),
		Price:       input.Price,
	}

	err := h.useCase.Update(ctx, domain)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, "")
}

func (h *Handler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := uuid.Parse(idParam)

	domain := &entity.Product{
		ID: id,
	}

	err := h.useCase.Delete(ctx, domain)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, "")
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	res, err := h.useCase.GetProducts(ctx)
	if err != nil {
		return
	}

	responseItems := make([]*model.ProductResponseDTO, 0, len(res.Result))

	for _, item := range res.Result {
		responseItems = append(responseItems, &model.ProductResponseDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Category:    model.ProductCategory(item.Category),
			Price:       item.Price,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	output := &model.ProductResponseList{
		Result: responseItems,
		Count:  res.Count,
	}

	ctx.JSON(http.StatusOK, output)
}
