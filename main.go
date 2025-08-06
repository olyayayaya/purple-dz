package main

import (
	"fmt"
	"math"
)

func main(){
const stepenVosv = 2 //untyped int сам переопределится при необходимости
height := 1.7
var ves float64 
ves = 60
//два вида записи переменных. если используется первый,  то там нельзя явно указать тип, это происходит автоматически
//вар используется в 2ух случ: когда нужно явно сконвертировать тип и когда не нужно сразу присваивать значение переменной
// var height, ves float64 = 1.7, 60
// height, ves := 1.7, 60

var calculate = ves / math.Pow(height, stepenVosv) //возведение в степень 2
fmt.Print(calculate)
}