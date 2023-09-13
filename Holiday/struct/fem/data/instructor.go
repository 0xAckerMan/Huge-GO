package data

import "fmt"

type Instructor struct {
    Id int
    Firstname string
    Lastname string
    Score int
}

func(i Instructor) Print() string{
    return fmt.Sprintf("%v %v (%d)", i.Firstname, i.Lastname, i.Id)
}

func NewInstructor(Fname string, Lname string) Instructor {
    return Instructor{Firstname: Fname, Lastname: Lname}
}
