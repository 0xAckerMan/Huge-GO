package main

import "fmt"

type Square struct {
    size int
    position [2]int
}

func (s *Square) SetSize(value int) {
    if value < 0 {
        panic ("size must be >= 0")
    }
    s.size = value
}

func (s *Square) GetSize() int {
    return s.size
}

func (s *Square) SetPosition(value [2]int) {
    if value[0] < 0 || value[1] < 0 {
        panic("position must be a tuple of 2 positive integers")
    }

    s.position = value
}

func NewSquare(size int, position [2]int) *Square {
    s := &Square{}
    s.SetSize(size)
    s.SetPosition(position)
    return s
}

func (s *Square) Area() int {
    return s.size * s.size
}

func (s *Square) My_Print() {
    if s.size == 0 {
        fmt.Println()
    }

    for y:= 0; y < s.position[1]; y++ {
        fmt.Println()
    }

    for i := 0; i < s.size; i++ {
        for x:=0; x<s.position[0]; x++ {
            fmt.Print(" ")
        }

        for j:=0; j < s.size; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}

func main () {
    square := NewSquare(4, [2]int{0, 0})
    square.My_Print()
}
