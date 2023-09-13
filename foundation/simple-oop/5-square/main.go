package main

import "fmt"

type Square struct {
    size int
}

func (s *Square) SetSize(value int) {
    if value < 0 {
        panic ("Size must be >= 0")
    }
    s.size = value
}

func (s *Square) GetSize() int {
    return s.size
}

func NewSquare (size int) *Square{
    s := &Square{}
    s.SetSize(size)
    return s
}

func (s *Square) Area() int {
    return s.size * s.size
}

func (s *Square) myprint() {
    for i := 0; i < s.size; i++ {
        for j := 0; j < s.size; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}


func main () {
    square := NewSquare(5)
    square.myprint()
}
