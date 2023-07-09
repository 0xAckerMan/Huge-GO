package main

import "fmt"

func print_most_numbers(){
    for i := 0; i <= 9; i++{
        if i != 2 && i != 4 {
            fmt.Print(i)
        }
    }
    fmt.Println()
}
