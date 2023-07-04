package x00printalphabet

import "fmt"


func IncrementAlpha(){
    for i := 'a'; i < 'z'; i++{
        fmt.Printf("%c ", i)
    }
}
