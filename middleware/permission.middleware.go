package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-assessment/entity"
	response2 "ordent-assessment/response"
)

func PermissionMiddleware(allowPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet("currentUser").(entity.User)
		userHasPermission := false
		for _, permission := range currentUser.Permissions {
			if permission == allowPermission {
				userHasPermission = true
			}
		}

		if !userHasPermission {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}
	}
}
