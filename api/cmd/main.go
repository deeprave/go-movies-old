package main

import (
	"flag"
	"fmt"
	"github.com/chigopher/pathlib"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

const version = "0.9.0"

func main() {
	cfg := AppConfig{
		Version: version,
	}

	flag.IntVar(&cfg.Port, "port", 4000, "listen port")
	flag.StringVar(&cfg.Host, "host", "localhost", "listen host")
	flag.StringVar(&cfg.Env, "env", "dev", "application environment (dev|staging|prod)")
	showVersion := flag.Bool("version", false, "program version")

	flag.Parse()

	if *showVersion {
		exe := pathlib.NewPath(os.Args[0])
		fmt.Printf("%s v%s\n", exe.Name(), version)
		os.Exit(0)
	}

	app := NewApplication(cfg, os.Stdout, "main ")
	startServer(app)
}

func startServer(app *Application) {
	address := fmt.Sprintf("%s:%d", app.Config.Host, app.Config.Port)

	var router *httprouter.Router
	// create the router
	router = httprouter.New()
	// allow for middleware here
	// add our application routes
	router = app.addStatusHandler(router)
	router = app.addMoviesHandlers(router)

	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Println("Listening on", address)

	err := srv.ListenAndServe()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
