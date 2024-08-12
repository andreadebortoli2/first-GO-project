package main

import "net/http"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl.html")
}
