package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zubroide/go-api-boilerplate/logger"
	"github.com/zubroide/go-api-boilerplate/model/service"
)

type UserController struct {
	service service.UserServiceInterface
	logger  logger.LoggerInterface
}

type UserListParameters struct {
	Name string
}

func NewUserController(service service.UserServiceInterface, logger logger.LoggerInterface) *UserController {
	return &UserController{service, logger}
}

func (c *UserController) List(ctx *gin.Context) {
	var params UserListParameters
	ctx.ShouldBindQuery(&params)
	list := c.service.GetUsers(params.Name)
	ctx.JSON(200, list)
}
