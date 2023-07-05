package main

import (
	"fmt"

	"go.dev/data"
)

var message string = "Go is funnn"

func calculateTax(price float32) (float32, float32) {
    return price*0.09, price*0.02
}

func main(){
    stateTax, cityTax := calculateTax(100)
    fmt.Println(stateTax, cityTax)
    printData()
    data.Run()
}
