package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-assessment/config"
	"ordent-assessment/entity"
	"ordent-assessment/request/auth_request"
	"ordent-assessment/response"
	"ordent-assessment/service"
	"strings"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authController struct {
	configuration    config.Config
	authService      service.AuthService
	userTokenService service.UserTokenService
}

func NewAuthController(authService service.AuthService, configuration config.Config, userTokenService service.UserTokenService) *authController {
	return &authController{authService: authService, configuration: configuration, userTokenService: userTokenService}
}

// Register User Registration godoc
// @Summary Register user.
// @Description Register user.
// @Tags Authentication
// @Accept */*
// @Param register body auth_request.RegisterRequest true "register user"
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /register [post]
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

// Login User Login godoc
// @Summary Login user.
// @Description Login user.
// @Tags Authentication
// @Accept */*
// @Param login body auth_request.LoginRequest true "login user"
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /login [post]
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

// Logout User Logout godoc
// @Summary Logout user.
// @Description Logout user.
// @Tags Authentication
// @Accept */*
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /logout [post]
// @Security BearerAuth
func (controller *authController) Logout(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.User)

	token := getTokenFromHeader(ctx)
	err := controller.userTokenService.Delete(ctx, currentUser.Id.Hex(), token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "logout success",
		Data:    nil,
	})
	return
}

func getTokenFromHeader(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	return tokenString
}
