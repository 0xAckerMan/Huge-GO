package main

import "joelkores.com/go/io/data"
import "fmt"

var globalVar = "This is a global var"
func main() {
    message := "hello from a module"
    //    print("hello world from a module")
    fmt.Println(message, globalVar)

    printData()
    data.Collection()
}
