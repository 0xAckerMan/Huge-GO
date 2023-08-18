package main

import "fmt"

func main() {
    my_list := int[] {1, 2, 3, 4, 5}
    idx := 3
    new_element := 9
    new_list := new_in_list(my_list, idx, new_element)
    fmt.Println(new_list)
}
