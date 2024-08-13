package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andreadebortoli2/first-GO-project/pkg/config"
	"github.com/andreadebortoli2/first-GO-project/pkg/handlers"
	"github.com/andreadebortoli2/first-GO-project/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	// set built or dev mode
	app.UseCache = false

	// set the config as repo for the handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache
	// set the config template cache available to render package
	render.NewTemplates(&app)

	/*
		// defining routes, using them and serve all in main(not best practice)
		http.HandleFunc("/", handlers.Repo.Home)
		http.HandleFunc("/about", handlers.Repo.About)
		_ = http.ListenAndServe(portNumber, nil) */

	_, _ = fmt.Printf("Starting application on port %s \n", portNumber)

	// routing with pat
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
