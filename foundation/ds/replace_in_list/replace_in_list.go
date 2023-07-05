package main

func replace_in_list(mylist []int, idx int, element int) []int {
    if idx < 0 || idx > len(mylist){
        return mylist
    }

    mylist[idx] = element

    return mylist

}
