package test_route

import (
	"github.com/gin-gonic/gin"
	test_route_controller "hospital-rest/controllers/test_route"
)

func Setup_test_route() *gin.Engine {
	router := gin.Default()

	test_route := router.Group("/test")
	{
		test_route.GET("", test_route_controller.Do_something)
	}
	return router
}
