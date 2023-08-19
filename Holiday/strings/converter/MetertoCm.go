package main

import "fmt"


type cm int
type m int

func (length m)  MeterstoCm() cm{
    return cm(length * 100) 
}

func ( length cm) CMtoMeters() m {
    return m(length / 100)
}

func test(){
    meters := m(4)
    incm := meters.MeterstoCm()
    Centimeters := cm(1000)
    inm := Centimeters.CMtoMeters()
    fmt.Printf("%d meters\n", inm)
    fmt.Printf("%d cm\n", incm)
}
