package main

import "fmt"

func calculator() {
    var operation string
    var int1, int2 int

    fmt.Println("Welcome to my simple GO calculator")
    fmt.Println("==================================")
    fmt.Println("Enter the operation: add, mul, div, sub")

    fmt.Scanf("%s", &operation)
    // fmt.Println(operation)

    fmt.Println("Enter the first number:")
    fmt.Scanf("%d", &int1)
    fmt.Println("Enter the second number:")
    fmt.Scanf("%d", &int2)

    switch {
        case operation == "add":
            fmt.Printf("the answer is %d\n ", int1 + int2)
        case operation == "mul":
            fmt.Printf("The answer is %d\n ", int1 * int2)
        case operation == "div":
            fmt.Printf("The answer is %d\n ", int1 / int2)
        case operation == "sub":
            fmt.Printf("The answer is %d\n ", int1 - int2)

    }
}
