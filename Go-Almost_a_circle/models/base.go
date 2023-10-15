package models

type Base struct{
    Id int
}

var nb_objects int

func NewBase(id int) *Base{
    if id >= 1 {
        return &Base{Id: id}
    }

    nb_objects ++

    return &Base{Id: nb_objects}
}
