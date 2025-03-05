package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessJSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func BadRequestJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func ServerErrorJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
}
