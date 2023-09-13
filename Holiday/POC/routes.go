package main

import "net/http"

func (app *Application) routes()*http.ServeMux{
    mux := http.NewServeMux()
    mux.HandleFunc("/api/healthcheck", app.healthcheck)
    return mux
}
