package main

import (
	"fmt"
	"simple-fem/backend/data"
)

func main(){
    max := data.Instructor{Id: 1, Firstname: "maximilliano", Lastname: "Firtman"}
    kyle := data.NewInstructor("Kyle", "Simsons")
    fmt.Println(max.Print())
    fmt.Println(kyle.Print())

    Node := data.Course{Name: "Node Class", Teacher: max, Duration: "4 hours"}
    fmt.Println(Node)
}
