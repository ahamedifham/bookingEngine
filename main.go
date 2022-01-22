package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"


func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2,3)
	_,_ = fmt.Fprintf(w, "This is the about page and 2+3 is %d", sum)
}

func addValues(x, y int) int {
	var sum int
	sum = x+y
	return sum
}

func main() {
	// fmt.Println("Hello world")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}