package main

import "fmt"


func print_list_integer(my_list[] int) {
    for i:=1; i < len(my_list); i++{
        fmt.Println(my_list[i])
    }
}
