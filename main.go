package main

import (
	"github.com/gin-gonic/gin"
	"ordent-assessment/config"
	"ordent-assessment/controller"
	"ordent-assessment/repository"
	"ordent-assessment/route"
	"ordent-assessment/service"
)

func main() {

	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "your app is healthy",
		})
	})

	api := r.Group("api/v1")

	route.ProductRoute(api, productController)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
