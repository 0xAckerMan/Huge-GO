package main

import (
    "fmt"
    "net/http"
)

func (app *Aplication) healthcheck (w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Status: Active")
    fmt.Fprintf(w, "Environment: %v\n", app.env)
    fmt.Fprintf(w, "Version: %v\n", version)
}

