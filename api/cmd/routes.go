package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) Routes(router *httprouter.Router) http.Handler {

	return router
}
