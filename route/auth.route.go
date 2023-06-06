package route

import (
	"github.com/gin-gonic/gin"
	"ordent-assessment/controller"
)

func AuthRoute(route *gin.RouterGroup, controller controller.AuthController, authMiddleware gin.HandlerFunc) {
	route.POST("/register", controller.Register)
	route.POST("/login", controller.Login)
	route.POST("/logout", authMiddleware, controller.Logout)
}
