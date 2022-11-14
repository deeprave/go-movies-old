package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *Application) statusHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.Config.Env,
		Version:     app.Config.Version,
	}
	if res, err := json.Marshal(currentStatus); err != nil {
		app.Logger.Println("json.Marshal", err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//goland:noinspection GoUnhandledErrorResult
		if _, err = w.Write(res); err != nil {
			app.Logger.Println("web.Write", err)
		}
	}
}

func (app *Application) addStatusHandler(router *httprouter.Router) *httprouter.Router {
	router.GET("/status", app.statusHandler)
	return router
}
