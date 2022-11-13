package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/dorukbulut/bookings/pkg/config"
	"github.com/dorukbulut/bookings/pkg/handlers"
	"github.com/dorukbulut/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main app function
func main() {

	//change this thing to true when in production
	app.InProduction = false

	//generate new session for web app.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s \n", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
