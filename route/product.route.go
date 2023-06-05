package route

import (
	"github.com/gin-gonic/gin"
	"ordent-assessment/controller"
)

func ProductRoute(route *gin.RouterGroup, controller controller.ProductController) {
	route.GET("/products", controller.GetAll)
	route.POST("/products", controller.Create)
	route.GET("/products/:id", controller.GetById)
	route.PUT("/products/:id", controller.Update)
	route.DELETE("/products/:id", controller.Delete)
}
