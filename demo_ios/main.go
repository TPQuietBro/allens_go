package main

import (
	"demo_ios/controller"
	"net/http"
)

func main() {
	controller.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
