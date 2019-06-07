package controller

import (
	"go-api-boilerplate/logger"
	"github.com/gin-gonic/gin"
)

type BaseControllerInterface interface {
}

type BaseController struct {
	logger *logger.Logger
}

func NewBaseController(logger *logger.Logger) *BaseController {
	return &BaseController{logger: logger}
}

func (c BaseController) response(context *gin.Context, obj interface{}, code int) {
	switch context.GetHeader("Accept") {
		case "application/xml":
			context.XML(code, obj)
		default:
			context.JSON(code, obj)
	}
}
