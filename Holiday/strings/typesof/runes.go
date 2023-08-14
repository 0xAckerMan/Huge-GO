package main

import (
	"fmt"
)

func runes(){
    s := "élite"

    fmt.Printf("%8T %[1]v\n", s)
    fmt.Printf("%8T %[1]v\n",[]rune(s))
    fmt.Printf("%8T %[1]v\n", []byte(s))
}
