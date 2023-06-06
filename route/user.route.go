package route

import (
	"github.com/gin-gonic/gin"
	"ordent-assessment/controller"
)

func UserRoute(route *gin.RouterGroup, controller controller.UserController) {
	route.GET("/users", controller.GetAll)
}
