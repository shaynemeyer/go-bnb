package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/shaynemeyer/go-bnb/internal/config"
	"github.com/shaynemeyer/go-bnb/internal/handlers"
	"github.com/shaynemeyer/go-bnb/internal/helpers"
	"github.com/shaynemeyer/go-bnb/internal/models"
	"github.com/shaynemeyer/go-bnb/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// Home is the home page handler

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is running on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when running in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
