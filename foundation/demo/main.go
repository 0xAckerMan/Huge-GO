package main

import "fmt"


func main(){
//     fmt.Println("Hello from go")
//     fmt.Println(split(6))
    price := 230
    priceInc(&price)
    fmt.Println(price)
// }
//
//
// func split(num int) (int, int) {
//     half := num/2
//     squared := num * 2
//
//     return half, squared

}


// func init(){

func priceInc(price *int) {
    *price ++
}
