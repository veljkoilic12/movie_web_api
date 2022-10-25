package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// declare a string containing app version number as a hard-coded, global constant (for now)
const version = "1.0.0"

// define a config struct to hold all the configuration settings for this app
// we read these settings from command-line flags when the app starts
type config struct {
	port int
	env  string
}

// define an application structure to hold the dependencies for our http handlers, helpers and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// declare an instance of the config struct
	var cfg config

	// read the values of the port and env command-line flags into the config struct
	// we also make default values for our flags in case if no corresponding flags are provided
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production")
	flag.Parse()

	// initialize a new logger which writes messages to the standard output stream, prefixed with
	// current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// as we now have our cfg and logger, we can now declare an instance of app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare a httpserver with some sensible timeout settings which listens on the port
	// provided in the config struct and uses httprouter from routes.go we as the handler
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
