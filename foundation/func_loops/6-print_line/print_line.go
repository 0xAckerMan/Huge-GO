package main

import "fmt"

func print_line(n int) {
    if n <= 0 {
        fmt.Println()
    }
    for i := 0; i < n; i++ {
        fmt.Print("_")
    }
    fmt.Println()
}
