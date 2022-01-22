package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"


func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
