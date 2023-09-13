package data

import "fmt"

type Duration string

type Course struct{
    Id int
    Name string
    Duration Duration
    Teacher Instructor
}

func (c Course) String () string {
    return fmt.Sprintf("---%v----%v---", c.Name, c.Teacher.Firstname)
}
