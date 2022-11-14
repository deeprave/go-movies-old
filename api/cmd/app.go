package main

import (
	"io"
	"log"
)

type AppConfig struct {
	Port    int
	Host    string
	Env     string
	Version string
}

type Application struct {
	Config AppConfig
	Logger *log.Logger
}

func NewApplication(cfg AppConfig, logwriter io.Writer, prefix string) *Application {
	logger := log.New(logwriter, prefix, log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds)
	return &Application{
		Config: cfg,
		Logger: logger,
	}
}
