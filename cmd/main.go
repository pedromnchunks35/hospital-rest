package main

import (
	"fmt"
	test_route_routes "hospital-rest/routes/test_route"
)

func main() {
	test_route := test_route_routes.Setup_test_route()
	fmt.Println("Starting the server at :8080")
	err := test_route.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
