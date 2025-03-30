package routes

import (
	"hospital-rest/controllers"
	"hospital-rest/services"
)

func RegisterChaincodeRoutes(fabricService services.FabricService) {
	routeObject.Group("/chaincode")
	{
		chaincodeController := controllers.CreateChaincodeController(fabricService)
		routeObject.GET("/:page/:size", chaincodeController.GetAllAssets)
		routeObject.GET("/getById/:id", chaincodeController.GetById)

		routeObject.POST("", chaincodeController.PostAsset)

		routeObject.PATCH("/patchById/:id", chaincodeController.PatchById)

		routeObject.DELETE("/deleteById/:id", chaincodeController.DeleteById)
	}
}
