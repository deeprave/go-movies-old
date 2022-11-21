package main

import (
	"github.com/julienschmidt/httprouter"
	"go-movies/api/models"
	"net/http"
	"strconv"
	"time"
)

func (api *Api) GetStatusHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type AppStatus struct {
		Status      string `json:"status"`
		Environment string `json:"environment"`
		Version     string `json:"version"`
	}
	cfg := api.Config()
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: cfg.Env,
		Version:     cfg.Version,
	}
	_, err := api.app.ModelToJson(w, http.StatusOK, currentStatus, "status")
	if err != nil {
		api.Log("status: ", err)
	}
}

// temporary, this is get movies from db...
func getMovie(id int) *models.Movie {
	return &models.Movie{
		Id:          id,
		Title:       "Placeholder title",
		Description: "Placeholder description",
		Year:        2022,
		ReleaseDate: time.Date(2022, time.April, 1, 0, 0, 0, 0, time.Local),
		Runtime:     150,
		Rating:      214,
		MPAARating:  "G",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func getMovies() []*models.Movie {
	movies := make([]*models.Movie, 0, 1)
	movie := getMovie(1)
	return append(movies, movie)
}

func (api *Api) GetMovieHandler(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")
	id, err := strconv.Atoi(idStr)
	var status = http.StatusBadRequest
	if err == nil {
		status, err = api.app.ModelToJson(w, http.StatusOK, getMovie(id), "movie")
	}
	if err != nil {
		api.Error(w, status, err)
	}
}

func (api *Api) GetMoviesHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	movies := getMovies()
	var (
		status = http.StatusOK
		err    error
	)
	if status, err = api.app.ModelToJson(w, status, movies, "movie"); err != nil {
		api.Error(w, status, err)
	}
}

func (api *Api) AddRoutes() {
	// allow for middleware here

	// add our application routes
	api.AddHandler(http.MethodGet, "/status", api.GetStatusHandler)
	api.AddHandler(http.MethodGet, "/movies/:id", api.GetMovieHandler)
	api.AddHandler(http.MethodGet, "/movies", api.GetMoviesHandler)
}
