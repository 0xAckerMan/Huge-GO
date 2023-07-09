package main

import "fmt"

func more_numbers() {
    for i := 0; i <= 9; i++ {
        for j := 0; j <= 14; j++ {
            fmt.Print(j)
        }
        fmt.Println()
    } 
    fmt.Println()
}
