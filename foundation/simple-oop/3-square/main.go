package main

import "fmt"

type Square struct{
    size int
}

func (s *Square) SetSize(value int){
    if value < 0 {
        panic("Size must be an integer")
    }
    s.size = value
}

func (s *Square) Area() int {
    return s.size * s.size
}

func main () {
    square := &Square{}
    square.SetSize(5)
    area := square.Area()

    fmt.Println("The size is: ", square.size)
    fmt.Println("The area is: ", area)
}
