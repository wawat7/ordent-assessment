package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"ordent-assessment/config"
	"ordent-assessment/controller"
	"ordent-assessment/database/seeder"
	_ "ordent-assessment/docs"
	"ordent-assessment/middleware"
	"ordent-assessment/repository"
	"ordent-assessment/route"
	"ordent-assessment/service"
)

// @title E-Commerce API
// @version 1.0
// @description This is a API E-Commerce.
// @termsOfService http://swagger.io/terms/

// @contact.name Wawat Prigala
// @contact.url https://wawatprigala.netlify.app
// @contact.email wawatprigala00@gmail.com

// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	runSeeder(database)

	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	userTokenRepository := repository.NewUserTokenRepository(database)
	userTokenService := service.NewUserTokenService(userTokenRepository)

	authService := service.NewAuthService(userService, userTokenService)
	authController := controller.NewAuthController(authService, configuration, userTokenService)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "your app is healthy",
		})
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFile("/doc.json", "./path/to/doc.json")
	api := r.Group("api/v1")

	route.ProductRoute(api, productController, middleware.AuthMiddleware(authService, userService, configuration, userTokenService))
	route.AuthRoute(api, authController, middleware.AuthMiddleware(authService, userService, configuration, userTokenService))
	route.UserRoute(api, userController)

	r.Run(":4000") // listen and serve on 0.0.0.0:4000

}

func runSeeder(db *mongo.Database) {
	seeder.UserSeeder(db)
}
