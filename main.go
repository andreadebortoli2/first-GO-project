package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the about page \n")

	x := 2
	y := 9
	sum := addValues(x, y)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("The sum of %d + %d is %d", x, y, sum))
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100.0
	var y float32 = 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprint(err))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot dive by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
