package main

import (
	"fmt"
	"os"
)

func SenR(){
    if len(os.Args) < 3 {
        fmt.Fprintln(os.Stderr, "Not enough args")
        os.Exit(-1)
    }
}
