package product_request

type UpdateProductRequest struct {
	Name     string `json:"name" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Category string `json:"category" binding:"required"`
}
