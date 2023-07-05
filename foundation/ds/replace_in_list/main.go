package main

import "fmt"

func main(){
    mylist := []int {1, 2, 3, 4, 5}
    idx := 3
    element := 9
    fmt.Println(replace_in_list(mylist, idx, element))
}
