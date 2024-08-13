package main

import (
	"net/http"

	"github.com/andreadebortoli2/first-GO-project/pkg/config"
	"github.com/andreadebortoli2/first-GO-project/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes handle all the routes
func routes(app *config.AppConfig) http.Handler {
	/*
		// routing with pat
		mux := pat.New()
		mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handlers.Repo.About)) */

	// routing with chi
	mux := chi.NewRouter()

	// recover from panics, log the panic and return an HTTP 500 status(if possible)
	mux.Use(middleware.Recoverer)
	// middleware test func
	mux.Use(WriteToConsole)
	// add crsf token in cookies for navigation security
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
