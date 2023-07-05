package data

import "fmt"


var Countries [10]string
var Slice []int
var Codes map[int]string

func init(){
    Countries[0] = "Kenya"
    Countries[1] = "Japan"
    Countries[2] = "USA"


    fmt.Println("Countries saved")
}
