package main

import "fmt"


func main(){
    fmt.Println("Hello from go")
    fmt.Println(split(6))
}


func split(num int) (int, int) {
    half := num/2
    squared := num * 2

    return half, squared
}


func init(){
    fmt.Println("This is just init")
}
