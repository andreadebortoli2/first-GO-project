package main

import (
	"net/http"

	"github.com/andreadebortoli2/first-GO-project/pkg/config"
	"github.com/andreadebortoli2/first-GO-project/pkg/handlers"
	"github.com/bmizerany/pat"
)

// routes handle all the routes
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
