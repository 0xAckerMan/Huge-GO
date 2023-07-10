package main

import "os"

func fileReader(filename string) (string, error) {
    content, err := os.ReadFile(filename)
    if err != nil {
        //handle the error
        return "", err
    }else{
        //read ops successful
        return string(content), nil
    }
}
