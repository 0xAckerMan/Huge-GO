package main

import "fmt"

func print_diagonal (n int) {
    if n <= 0 {
        fmt.Println()
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(" ")
        }
        fmt.Println("\\")
    }
}
