package data

// import "fmt"


// func Run(){
//     fmt.Println(split(25))
// }
//
func Split(sum int) (int, int) {
    x := sum / 2
    y := sum - x
    return x, y
}
