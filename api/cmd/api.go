package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-movies/api/app"
	"log"
	"net/http"
	"os"
	"time"
)

type Api struct {
	app    *app.Application
	router *httprouter.Router
}

func NewApi(app *app.Application) *Api {
	return &Api{
		app:    app,
		router: httprouter.New(),
	}
}

func (api *Api) Address() string {
	cfg := api.app.Config
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func (api *Api) Logger() *log.Logger {
	return api.app.Logger
}

func (api *Api) Log(v ...any) {
	api.Logger().Print(v...)
}

func (api *Api) startServer() {
	api.AddRoutes()
	address := api.Address()
	srv := &http.Server{
		Addr:         address,
		Handler:      api.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	api.Log("Listening on ", address)

	err := srv.ListenAndServe()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func (api *Api) Respond(w http.ResponseWriter, data interface{}) *app.Error {
	err := api.app.ModelToJson(w, http.StatusOK, data, "data")
	if err != nil {
		api.Error(w, http.StatusBadRequest, err)
	}
	return err
}

func (api *Api) Error(w http.ResponseWriter, status int, err error) {
	api.Log(err, status)
	_ = api.app.ErrorToJson(w, status, err)
}

func (api *Api) AddHandler(method string, path string, handler httprouter.Handle) {
	api.router.Handle(method, path, handler)
}
