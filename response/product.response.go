package response

import (
	"ordent-assessment/entity"
	"time"
)

type productResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatProduct(product entity.Product) *productResponse {
	return &productResponse{
		Id:        product.Id.Hex(),
		Name:      product.Name,
		Price:     product.Price,
		Category:  product.Category,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func FormatProducts(products []entity.Product) []*productResponse {
	formats := []*productResponse{}

	for _, product := range products {
		format := &productResponse{
			Id:        product.Id.Hex(),
			Name:      product.Name,
			Price:     product.Price,
			Category:  product.Category,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}

		formats = append(formats, format)
	}

	return formats
}
