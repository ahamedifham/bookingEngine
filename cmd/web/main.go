package main

import (
	"fmt"
	"net/http"
	"github.com/ahamedifham/bookingEngine/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
