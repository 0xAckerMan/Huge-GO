package main

import (
	"fmt"
	"net/http"
)

// Asimple http server

func MyServer(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "Hello from go")
}

