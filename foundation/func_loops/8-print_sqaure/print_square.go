package main

import "fmt"

func print_square(size int) {
    if size <= 0 {
        fmt.Println()
    }
    for i := 1; i <= size; i++{
        for j := 1; j < size; j++ {
            fmt.Print("#")
        }
        fmt.Println("#")
    }
}
