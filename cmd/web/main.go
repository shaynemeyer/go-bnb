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
	"github.com/shaynemeyer/go-bnb/internal/driver"
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
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Server is running on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// change this to true when running in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=shaynemeyer password=")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
