package main

import (
	"log"

	"net/http"

	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/firdisml/go-http/config"

	"github.com/firdisml/go-http/handlers"

	"github.com/firdisml/go-http/renderer"
)

var app config.AppConfig

var Session *scs.SessionManager

func main() {

	app.Production = false

	Session = scs.New()

	Session.Lifetime = 24 * time.Hour

	Session.Cookie.Persist = true

	Session.Cookie.SameSite = http.SameSiteLaxMode

	Session.Cookie.Secure = app.Production

	app.Session = Session

	tc, err := renderer.CacheTemplate()

	if err != nil {

		log.Fatal("cannot create tempalte cache")

	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	renderer.NewTemplates(&app)

	srv := &http.Server{
		Addr:    ":1337",
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {

		log.Fatal(err)

	}

}
