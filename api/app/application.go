package app

import (
	"encoding/json"
	"go-movies/api/helpers"
	"io"
	"log"
	"net/http"
)

type Error struct {
	message string
	status  int
}

// satisfy stdlib Error interface

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Status() int {
	if e.status == 0 {
		// default status
		return http.StatusBadRequest
	}
	return e.status
}

func NewError(message string, status int) *Error {
	err := Error{
		message: message,
		status:  status,
	}
	return &err
}

type Application struct {
	Config *helpers.AppConfig
	Logger *log.Logger
}

func NewApplication(cfg *helpers.AppConfig, logwriter io.Writer, prefix string) *Application {
	logger := log.New(logwriter, prefix, log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds)
	return &Application{
		Config: cfg,
		Logger: logger,
	}
}

func (app *Application) ModelToJson(w http.ResponseWriter, status int, data interface{}, wrap string) *Error {
	var (
		jsonstr []byte
		err     error
	)
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	if jsonstr, err = json.Marshal(wrapper); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, err = w.Write(jsonstr)
	}
	if err != nil {
		return NewError(err.Error(), 0)
	}
	return nil
}

func (app *Application) ErrorToJson(w http.ResponseWriter, status int, error error) error {
	type ErrorReport struct {
		Message string `json:"message"`
	}
	return app.ModelToJson(w, status, ErrorReport{error.Error()}, "errors")
}
