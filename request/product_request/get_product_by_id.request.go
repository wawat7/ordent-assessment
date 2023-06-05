package product_request

type GetProductById struct {
	Id string `uri:"id" binding:"required"`
}
