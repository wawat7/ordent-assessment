package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-assessment/request/product_request"
	"ordent-assessment/response"
	"ordent-assessment/service"
)

type ProductController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type productController struct {
	service service.ProductService
}

func NewProductController(srv service.ProductService) *productController {
	return &productController{service: srv}
}

// GetAll Product godoc
// @Summary Show all products.
// @Description get all of products.
// @Tags Product
// @Accept */*
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /products [get]
// @Security BearerAuth
func (controller *productController) GetAll(ctx *gin.Context) {
	products, err := controller.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    response.FormatProducts(products),
	})
	return
}

// GetById Product godoc
// @Summary Get Product by ID.
// @Description Find product by ID.
// @Tags Product
// @Accept */*
// @Param productId path string true "Get product by ID"
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /products/{productId} [get]
// @Security BearerAuth
func (controller *productController) GetById(ctx *gin.Context) {
	var inputId product_request.GetProductById
	err := ctx.ShouldBindUri(&inputId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	product, err := controller.service.GetById(ctx, inputId.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    response.FormatProduct(product),
	})
	return
}

// Create Product godoc
// @Summary Create product.
// @Description create new product.
// @Tags Product
// @Accept */*
// @Param product body product_request.CreateProductRequest true "create product"
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /products [post]
// @Security BearerAuth
func (controller *productController) Create(ctx *gin.Context) {
	var input product_request.CreateProductRequest
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	productId, err := controller.service.Create(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    gin.H{"id": productId},
	})
	return
}

// Update Product godoc
// @Summary Update product by ID.
// @Description Update product by ID.
// @Tags Product
// @Accept */*
// @Param productId path string true "Update product by ID"
// @Param product body product_request.UpdateProductRequest true "update product"
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /products/{productId} [put]
// @Security BearerAuth
func (controller *productController) Update(ctx *gin.Context) {
	var inputId product_request.GetProductById
	err := ctx.ShouldBindUri(&inputId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var input product_request.UpdateProductRequest
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = controller.service.UpdateById(ctx, inputId.Id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    nil,
	})
	return
}

// Delete Product godoc
// @Summary Remove product by ID.
// @Description Remove product by ID.
// @Tags Product
// @Accept */*
// @Param productId path string true "Delete product by ID"
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /products/{productId} [delete]
// @Security BearerAuth
func (controller *productController) Delete(ctx *gin.Context) {
	var inputId product_request.GetProductById
	err := ctx.ShouldBindUri(&inputId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = controller.service.DeleteById(ctx, inputId.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    nil,
	})
	return
}
