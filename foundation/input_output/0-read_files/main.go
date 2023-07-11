package main

import (
	"fmt"
	"os"
)

func main() {
    filedir, _ := os.Getwd()
    filename := filedir + "/my_file_0.txt"
    c, err := readFile(filename)
    if err == nil{
       fmt.Println(c) 
   } else {
       fmt.Println(err)
   }
}
