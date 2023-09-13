package main

import "net/http"

func (app *Application) route() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/healthcheck", app.healthcheck)

    return mux
}
