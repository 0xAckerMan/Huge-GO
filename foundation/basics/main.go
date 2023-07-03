package main

import "joelkores.com/go/io/data"
import "joelkores.com/go/io/pointers"
import "fmt"

var globalVar = "This is a global var"
func main() {
    message := "hello from a module"
    //    print("hello world from a module")
    fmt.Println(message, globalVar)

    printData()
    fmt.Println(data.Split(25))
    data.Collection()
    age := 20
    pointers.Birthday(&age)
    fmt.Println(age)
}
