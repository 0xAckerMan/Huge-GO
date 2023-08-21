package main

import (
	"fmt"
	"net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    err := http.ListenAndServe(":4000", mux)

    if err != nil {
        fmt.Println(err)
    }
}

func home (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, http.StatusText(
            http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
    }

    fmt.Fprintf(w, "Welcome Home")
}
