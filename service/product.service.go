package service

import (
	"context"
	"ordent-assessment/entity"
	"ordent-assessment/repository"
	"ordent-assessment/request/product_request"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]entity.Product, error)
	GetById(ctx context.Context, Id string) (entity.Product, error)
	Create(ctx context.Context, payload product_request.CreateProductRequest) (productId string, err error)
	UpdateById(ctx context.Context, Id string, payload product_request.UpdateProductRequest) error
	DeleteById(ctx context.Context, Id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService {
	return &productService{repo: repo}
}

func (service *productService) GetAll(ctx context.Context) ([]entity.Product, error) {
	products, err := service.repo.FindAll(ctx)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (service *productService) GetById(ctx context.Context, Id string) (entity.Product, error) {
	product, err := service.repo.FindById(ctx, Id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (service *productService) Create(ctx context.Context, payload product_request.CreateProductRequest) (productId string, err error) {
	product := entity.Product{
		Name:     payload.Name,
		Price:    payload.Price,
		Category: payload.Category,
	}

	productId, err = service.repo.Create(ctx, product)
	if err != nil {
		return productId, err
	}
	return productId, nil
}

func (service *productService) UpdateById(ctx context.Context, Id string, payload product_request.UpdateProductRequest) error {
	product, err := service.GetById(ctx, Id)
	if err != nil {
		return err
	}

	product.Name = payload.Name
	product.Price = payload.Price
	product.Category = payload.Category

	err = service.repo.Update(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (service *productService) DeleteById(ctx context.Context, Id string) error {
	product, err := service.GetById(ctx, Id)
	if err != nil {
		return err
	}

	err = service.repo.Delete(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
