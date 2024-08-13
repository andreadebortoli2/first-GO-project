package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/andreadebortoli2/first-GO-project/pkg/config"
	"github.com/andreadebortoli2/first-GO-project/pkg/handlers"
	"github.com/andreadebortoli2/first-GO-project/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change to treu when in production
	app.InProduction = false

	// define the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour              // how long a session can stay alive
	session.Cookie.Persist = true                  // set the sesssion persist in the coockie after the browser window close
	session.Cookie.SameSite = http.SameSiteLaxMode // default in go
	session.Cookie.Secure = app.InProduction       // encryption for coockies as required in https, set false in dev mode, true in production
	// set session as AppConfig session
	app.Session = session

	// generate template cache and set it to AppConfig template cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache

	// set production or dev mode
	app.UseCache = false

	// set the config as repo for the handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

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
