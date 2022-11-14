package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"go-movies/api/models"
	"net/http"
	"strconv"
)

func (app *Application) getMovieHandler(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	if id, err := strconv.Atoi(p.ByName("id")); err != nil {
		app.Logger.Print(errors.New("invalid id specified"))
	} else {
		movie := models.Movie{
			Id: id,
			Title: "Placeholder",
			Description: ""
			Year        int          `json:"year"`
			ReleaseDate time.Time    `json:"release_date"`
			Runtime     int          `json:"runtime"`
			Rating      int          `json:"rating"`
			MPAARating  string       `json:"mpaa_rating"`
			CreatedAt   time.Time    `json:"created_at"`
			UpdatedAt   time.Time    `json:"updated_at"`
			MovieGenre  []MovieGenre `json:"-"`

		}
	}
}

func (app *Application) getMoviesHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

}

func (app *Application) addMoviesHandlers(router *httprouter.Router) *httprouter.Router {
	router.GET("/movies/:id", app.getMovieHandler)
	router.GET("/movies", app.getMoviesHandler)
	return router
}
