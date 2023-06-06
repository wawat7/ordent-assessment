package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordent-assessment/response"
	"ordent-assessment/service"
)

type UserController interface {
	GetAll(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(srv service.UserService) *userController {
	return &userController{service: srv}
}

// GetAll User godoc
// @Summary Show all users.
// @Description get all of users.
// @Tags User
// @Accept */*
// @Produce application/json
// @Success 200 {object} response.BaseResponse
// @Router /users [get]
func (controller *userController) GetAll(ctx *gin.Context) {
	users, err := controller.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    response.FormatUsers(users),
	})
	return
}
