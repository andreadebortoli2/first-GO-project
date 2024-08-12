package main

import (
	"fmt"
	"net/http"

	"github.com/andreadebortoli2/first-GO-project/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
