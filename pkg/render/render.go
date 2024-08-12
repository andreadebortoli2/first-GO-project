package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/andreadebortoli2/first-GO-project/pkg/config"
	"github.com/andreadebortoli2/first-GO-project/pkg/models"
)

var app *config.AppConfig

// NewTemplates set config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// ? prepared for later usage
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplates renders templates using html/template
func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	/*
		// create a template cache
		templateCache, err := CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		} */

	var templateCache map[string]*template.Template

	// set if statement to work in built or dev mode
	if app.UseCache {
		// built mode
		// instead of create the template cache each time, i want to use the app config template cache
		templateCache = app.TemplateCache
	} else {
		// dev mode (re-create the template cache each time to check changes i make developing)
		templateCache, _ = CreateTemplateCache()
	}

	// get requested from cache
	requestedTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// ? preapared for later usage
	td = AddDefaultData(td)

	// if an error occur, for a precise check, i add a buffer to suspend the actions, execute the code and see where the error is coming from
	buf := new(bytes.Buffer)
	err := requestedTemplate.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	// if the buffer doesn't give back an error i can write it in the responseWriter to send back the page
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// when i render a page i first need the page template to render nd then all layout and partials so:
	// get all of the files that end with .page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all pages
	for _, page := range pages {
		// i save the final part of the path as name to be used as key in the map
		name := filepath.Base(page)
		// i save the parsed file with his name
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// i get all the layout files from ./templates to check later if there are layout that i need
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// if i have layout files (i need them) i parse them all
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// at the end of the loop i save in the myCache slice the parsed page + parsed layout and eventual parsed partials with the page file name as key
		myCache[name] = templateSet
	}

	// last i return the full myCache slice, where at every page name (key) i have the relevant parsed page followed by parsed layouts and partials (in the right position to be read by the render function)
	return myCache, nil
}

/* func RenderTemplates(w http.ResponseWriter, tmpl string) {
	// basic page rendering, need to read the disk any time a page is re-visited, not
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
} */

/*
// template cache
var tc = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template ot cache
	tc[t] = tmpl

	return nil
} */
