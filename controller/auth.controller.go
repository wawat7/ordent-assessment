package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-assessment/config"
	"ordent-assessment/request/auth_request"
	"ordent-assessment/response"
	"ordent-assessment/service"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
	configuration config.Config
	authService   service.AuthService
}

func NewAuthController(authService service.AuthService, configuration config.Config) *authController {
	return &authController{authService: authService, configuration: configuration}
}

func (controller *authController) Register(ctx *gin.Context) {
	var input auth_request.RegisterRequest
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = controller.authService.Register(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "register success",
		Data:    nil,
	})
	return
}

func (controller *authController) Login(ctx *gin.Context) {
	var input auth_request.LoginRequest
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	token, err := controller.authService.Login(ctx, input, controller.configuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data: gin.H{
			"access_token": token,
		},
	})
	return
}

func (controller *authController) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
