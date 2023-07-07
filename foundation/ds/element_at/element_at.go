package main

func element_at(mylist[]int, idx int) *int {
    if idx < 0 || idx >= len(mylist) {
        return nil
    }
    return &mylist[idx]
}
