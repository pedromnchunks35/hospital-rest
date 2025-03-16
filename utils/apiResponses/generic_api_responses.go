package apiResponses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadArgumentsResponse(context *gin.Context, message string) {
	context.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"code":    http.StatusBadRequest,
	})
}

func SuccessResponse(context *gin.Context, result string, message string) {
	context.JSON(http.StatusOK, gin.H{
		"result":  result,
		"code":    http.StatusOK,
		"message": message,
	})
}
