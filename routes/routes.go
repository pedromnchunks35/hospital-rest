package routes

import "github.com/gin-gonic/gin"

var routeObject = gin.Default()

func GetRouteObject() *gin.Engine {
	return routeObject
}
