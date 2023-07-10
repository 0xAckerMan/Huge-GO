package main

import "os"

func writeToFile(filename string, content string) error {
    return os.WriteFile(filename, []byte(content), 0644)
}
