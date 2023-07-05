package main

import "fmt"

func print_reversed_list_integer(my_list []int){
    for i := len(my_list)-1; i > 0; i-- {
        fmt.Println(my_list[i])
    }
}
