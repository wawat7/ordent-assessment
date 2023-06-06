package route

import (
	"github.com/gin-gonic/gin"
	"ordent-assessment/constant"
	"ordent-assessment/controller"
	"ordent-assessment/middleware"
)

func ProductRoute(route *gin.RouterGroup, controller controller.ProductController, authMiddleware gin.HandlerFunc) {
	route.GET("/products", authMiddleware, middleware.PermissionMiddleware(constant.CAN_READ_PRODUCT), controller.GetAll)
	route.POST("/products", authMiddleware, middleware.PermissionMiddleware(constant.CAN_CREATE_PRODUCT), controller.Create)
	route.GET("/products/:id", authMiddleware, middleware.PermissionMiddleware(constant.CAN_READ_PRODUCT), controller.GetById)
	route.PUT("/products/:id", authMiddleware, middleware.PermissionMiddleware(constant.CAN_UPDATE_PRODUCT), controller.Update)
	route.DELETE("/products/:id", authMiddleware, middleware.PermissionMiddleware(constant.CAN_DELETE_PRODUCT), controller.Delete)
}
