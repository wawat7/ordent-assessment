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
