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

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)

	userTokenRepository := repository.NewUserTokenRepository(database)
	userTokenService := service.NewUserTokenService(userTokenRepository)

	authService := service.NewAuthService(userService, userTokenService)
	authController := controller.NewAuthController(authService, configuration)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "your app is healthy",
		})
	})
	api := r.Group("api/v1")

	route.ProductRoute(api, productController)
	route.AuthRoute(api, authController)

	r.Run(":4000") // listen and serve on 0.0.0.0:8080

}
