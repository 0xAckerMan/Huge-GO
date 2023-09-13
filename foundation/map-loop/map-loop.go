package main

import "fmt"

var myMap = map[string]int {"http":80, "https": 443, "smb": 445}
func MapLoop(){

    for k := range myMap {
        fmt.Println(k, myMap[k])
    }
}

func KeyVal(){
    for k, v := range myMap {
        fmt.Println(k, v)
    }
}
