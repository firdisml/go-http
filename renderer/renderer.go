package renderer

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/firdisml/go-http/config"
	"github.com/firdisml/go-http/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {

		tc = app.TemplateCache

	} else {
		tc, _ = CacheTemplate()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}

	err := t.Execute(w, td)
	if err != nil {
		log.Fatal(err)
	}

}

// More Complex Style
func CacheTemplate() (map[string]*template.Template, error) {
	var myCache = make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		ls, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(ls) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}

//Caching Template Style

/* var tc = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	_, exist := tc[t]

	if !exist {
		err = cacheTemplate(t)
		if err != nil {
			fmt.Println("Error Using Template")
		}
	} else {
		fmt.Println("Using Cached Template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Using Template")
	}
}

func cacheTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		fmt.Println("Error Using Template")
	}

	tc[t] = tmpl

	return nil

}
*/
