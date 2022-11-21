package main

import (
	"go-movies/api/app"
	"go-movies/api/test"
	"testing"
)

var testConfig = app.Config{
	Port:    9999,
	Host:    "example.com",
	Env:     "prod",
	Version: "1.0.0",
}

func getApplication() *app.Application {
	return app.NewApplication(&testConfig, test.NewTestLog(), "test ")
}

func getApi() *Api {
	return &Api{
		app:    getApplication(),
		router: nil,
	}
}

func TestApi_Config(t *testing.T) {
	//api := getApi()

}
