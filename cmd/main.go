package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"hospital-rest/routes"
	"hospital-rest/services"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fabricService := services.NewFabricServiceImpl()
	routes.RegisterChaincodeRoutes(fabricService)

	runObject := routes.GetRouteObject()
	fmt.Println("Starting the server at :8080")

	err = runObject.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
