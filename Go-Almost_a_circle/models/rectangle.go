package models

type Rectangle struct{
    width int
    height int
    x int
    y int
    id *Base
}

func NewRectangle(width, height, x, y, id int) *Rectangle{
    base := NewBase(id)
    return &Rectangle{
        id: base,
        width: width,
        height: height,
        x: x,
        y: y,
    }
}


func (r *Rectangle) Height() int{
    return r.height
}

func (r *Rectangle) Width() int{
    return r.width
}

