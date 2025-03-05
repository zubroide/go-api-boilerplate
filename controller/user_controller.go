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
	Name string `json:"name" form:"name"`
}

func NewUserController(service service.UserServiceInterface, logger logger.LoggerInterface) *UserController {
	return &UserController{service, logger}
}

func (c *UserController) List(ctx *gin.Context) {
	var params UserListParameters
	if err := ctx.ShouldBind(&params); err != nil {
		BadRequestJSON(ctx, err.Error())
		return
	}
	list, err := c.service.GetUsers(params.Name)
	if err != nil {
		ServerErrorJSON(ctx, "Something went wrong")
		return
	}
	SuccessJSON(ctx, list)
}
