package main

import (
	"fmt"
	"os"
)

func main() {
    rootpath, _ := os.Getwd()
    writepath := rootpath + "/writetest.txt"
    c, err := fileReader(rootpath + "/test.txt")
    if err == nil {
        writevalue := fmt.Sprintf("Original message: %v \n I am enjoying golang", c)
        err := writeToFile(writepath, writevalue)
        if err != nil {
            fmt.Printf("FAILED: %v", err)
        }
    }else {
        fmt.Printf("You have an error %v", err)
    }
}
