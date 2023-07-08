package main

func isupper(c string) int { 
    for i := 65; i <= 90; i++ {
        if c == string(i) {
            return (1)
        }
    }
    return (0)
}
