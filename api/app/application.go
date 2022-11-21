package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func NewApplication(cfg *Config, logwriter io.Writer, prefix string) *Application {
	logger := log.New(logwriter, prefix, log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds)
	return &Application{
		cfg: cfg,
		log: logger,
	}
}

func (cfg *Config) Address() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

type Application struct {
	cfg *Config
	log *log.Logger
}

func (app *Application) Address() string {
	return app.cfg.Address()
}

func (app *Application) Logger() *log.Logger {
	return app.log
}

func (app *Application) Config() *Config {
	return app.cfg
}

func (app *Application) Log(v ...any) {
	app.Logger().Print(v...)
}

func (app *Application) ModelToJson(w http.ResponseWriter, status int, data interface{}, wrap string) (int, error) {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	if jsonstr, err := json.Marshal(wrapper); err != nil {
		return http.StatusNotAcceptable, err
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if _, err = w.Write(jsonstr); err != nil {
			return http.StatusBadGateway, err
		}
	}
	return http.StatusOK, nil
}

func (app *Application) ErrorToJson(w http.ResponseWriter, status int, error error) (int, error) {
	type ErrorReport struct {
		Message string `json:"message"`
	}
	return app.ModelToJson(w, status, ErrorReport{error.Error()}, "error")
}
