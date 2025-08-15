package main

import (
	"errors"
	"fmt"
	"strings"
)

func main(){
valuta := map[string]float64 {
"usdInrub" : 79.70,
"usdIneur" : 0.86,
"rubInusd" : 0.013,
"rubIneur" : 0.011,
"eurInusd" : 1.17,
"eurInrub" : 92.91,
}

fmt.Println("Конвертер валют") 
	for {
		valuta1, money, valuta2, err := readingUsersField()
			if err != nil {
			fmt.Println(err)
			continue
		}
		raschet(valuta1, money, valuta2, valuta)
		break
	}

// const usdEur = 0.86
// const usdRub = 80.23


// var calculateRUB = usdRub / usdEur
// fmt.Print(calculateRUB)
// readingUsersField()
}

func readingUsersField() (string, float64, string, error) {
var valuta1 string
fmt.Print("Введите исходную валюту (usd, eur, rub):")
fmt.Scan(&valuta1)
valuta1 = strings.ToLower(valuta1)

var money float64
fmt.Print("Введите число:")
fmt.Scan(&money)
	if money <= 0 {
		return "_", 0, "_", errors.New("не указано число")
	}

var valuta2 string
switch valuta1 {
	case "usd":
	fmt.Print("Введите целевую валюту (eur, rub):")
	fmt.Scan(&valuta2)
	valuta2 = strings.ToLower(valuta2)
		if valuta2 == "usd" || (valuta2 != "eur" && valuta2 != "rub") {
		return "_", 0, "_", errors.New("указана неверная целевая валюта")
	}
	case "eur":
	fmt.Print("Введите целевую валюту (usd, rub):")
	fmt.Scan(&valuta2)
	valuta2 = strings.ToLower(valuta2)
		if valuta2 == "eur" || (valuta2 != "usd" && valuta2 != "rub") {
		return "_", 0, "_", errors.New("указана неверная целевая валюта")
	}
	case "rub":
	fmt.Print("Введите целевую валюту (usd, eur):")
	fmt.Scan(&valuta2)
	valuta2 = strings.ToLower(valuta2)
		if valuta2 == "rub" || (valuta2 != "usd" && valuta2 != "eur") {
		return "_", 0, "_", errors.New("указана неверная целевая валюта")
	}
}

return valuta1, money, valuta2, nil
}

func raschet(valuta1 string, money float64, valuta2 string, valuta map[string]float64) {
var result float64

value, ok := valuta[valuta1 + "In" + valuta2]

if ok {
result = money * value
}

// if valuta1 == "usd" && valuta2 == "rub" {
// 	result = money * 79.70
// } else if valuta1 == "usd" && valuta2 == "eur" {
// 	result = money * 0.86
// } else if valuta1 == "rub" && valuta2 == "usd" {
// 	result = money * 0.013
// } else if valuta1 == "rub" && valuta2 == "eur" {
// 	result = money * 0.011
// } else if valuta1 == "eur" && valuta2 == "usd" {
// 	result = money * 1.17
// } else if valuta1 == "eur" && valuta2 == "rub" {
// 	result = money * 92.91
// }

output := fmt.Sprintf("результат конвертации: %.3f", result)
fmt.Println(output)
}