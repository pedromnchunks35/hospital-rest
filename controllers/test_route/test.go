package test_route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Do_something(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}
