package main

import "os"

func write_file(filename string, content string) (int, error) {
    err := os.WriteFile(filename, []byte(content), 0644)
    if err != nil{
        return 0, err
    }else{
        return len(content), nil
    }
}
