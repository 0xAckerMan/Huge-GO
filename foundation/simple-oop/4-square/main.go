package main

import "fmt"

type Square struct {
    size int
}

func (s *Square) setSize(value int) {
    if value < 0 {
        panic("size must be >= 0")
    }
    s.size = value
}


func (s *Square) getSize() int {
    return s.size
}

func NewSquare(size int) *Square{
    s := &Square{}
    s.setSize(size)
    return s
}

func (s *Square) area() int {
    return s.size * s.size
}

func main () {
    square := NewSquare(5)
    fmt.Println("The size is: ", square.getSize())
    fmt.Println("The area is: ", square.area())
}
