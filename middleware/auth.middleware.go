package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-assessment/config"
	response2 "ordent-assessment/response"
	"ordent-assessment/service"
	"strings"
)

func AuthMiddleware(authService service.AuthService, userService service.UserService, configuration config.Config, userTokenService service.UserTokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString, configuration)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		userID := claim["user_id"].(string)

		userToken, err := userTokenService.FindToken(c, userID, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		if userToken.Token != tokenString {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		user, err := userService.GetById(c, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response2.BaseResponse{
				Message: "Unauthorized",
				Data:    nil,
			})
			return
		}

		c.Set("currentUser", user)
	}
}
