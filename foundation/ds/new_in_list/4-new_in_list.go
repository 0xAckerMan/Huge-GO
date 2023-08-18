package main

func new_in_list(my_list []int, idx int, element int ) []int {
    temp := my_list
    temp[idx] = element
    return temp
}
