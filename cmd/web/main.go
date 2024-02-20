package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	filesServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", filesServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	logger.Info("starting server", slog.String("addr", ":4000"))

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
