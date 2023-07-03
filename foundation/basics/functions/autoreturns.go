package functions


func taxes(price float32) (stateTax, cityTax float32) {
    stateTax := price * 0.09
    cityTax := price * 0.015
    return
}

func main(){
    tax1, tax2 := taxes(100)
}
