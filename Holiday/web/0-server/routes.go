package main

import "net/http"

func (app *Aplication) routes () *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/healthcheck", app.healthcheck)
    return mux
}
