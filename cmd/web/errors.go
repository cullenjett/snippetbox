package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func (app *App) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *App) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
