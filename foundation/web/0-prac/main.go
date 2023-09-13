package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", MyServer)
    log.Fatal(http.ListenAndServe(":2000",nil))
}
