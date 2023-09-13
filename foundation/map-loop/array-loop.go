package main

import "fmt"

func arrayloop(){
    myarr := [5]string {"Joel", "Kor3s", "Kashu"}

    for _, i := range myarr {
        fmt.Println(i)
    }
}
