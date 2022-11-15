package main

import (
	"flag"
	"fmt"
	"github.com/chigopher/pathlib"
	"go-movies/api/app"
	"go-movies/api/helpers"
	"os"
)

const version = "0.9.0"

func main() {
	cfg := helpers.AppConfig{
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

	NewApi(app.NewApplication(&cfg, os.Stdout, "main ")).startServer()
}
