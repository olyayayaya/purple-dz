package main

import (
	"fmt"
)

func main(){
const usdEur = 0.86
const usdRub = 80.23


var calculateRUB = usdRub / usdEur
fmt.Print(calculateRUB)
readingUsersField()
}

func readingUsersField() float64 {
var valuta float64
fmt.Print("Введите значениие:")
fmt.Scan(&valuta)
return valuta
}

func raschet(number int, valuta1 string, valuta2 string) {

}