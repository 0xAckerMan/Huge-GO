package main

import (
	"fmt"
	"net/http"

)

func(app *Application) healthcheck(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Status: Active\n")
    fmt.Fprintf(w, "Environment: %s\n", app.env)
    fmt.Fprintf(w, "Status: %s\n", Version)
}
