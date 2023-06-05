package product_request

type CreateProductRequest struct {
	Name     string `json:"name" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Category string `json:"category" binding:"required"`
}
