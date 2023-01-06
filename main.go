package main

import (
	"booking-app/rest"
	"fmt"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	rest.HandleRequests()
}
