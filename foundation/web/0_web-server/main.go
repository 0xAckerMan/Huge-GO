package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w , "It works")
}

func main(){
    http.HandleFunc("/", home)
    fmt.Println("Running at port 3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
