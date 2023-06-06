package route

import (
	"github.com/gin-gonic/gin"
	"ordent-assessment/controller"
)

func AuthRoute(route *gin.RouterGroup, controller controller.AuthController) {
	route.POST("/register", controller.Register)
	route.POST("/login", controller.Login)
}
