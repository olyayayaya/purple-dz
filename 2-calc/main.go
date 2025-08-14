package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("operations") 
	operation, numbers, err := readingUsersField()

	if err != nil {
	fmt.Println(err)
	}

	result := calculation(operation, numbers)
	fmt.Println("результат:", result)

}

// func readingUsersField() () {
// var operation string
// fmt.Print("Введите операцию (avg - среднее, sum - сумма, med - медиана):")
// fmt.Scan(&operation)
// operation = strings.ToLower(operation)

// var strNumbers string
// // arrNumbers := []float64{}

// fmt.Print("Введите числа:")
// // fmt.Scanln(&strNumbers)
//     reader := bufi.NewReader(os.Stdin)
//     strNumbers, _ := reader.ReadString('\n')
//     strNumbers = strings.TrimSpace(strNumbers)
// // strNumbersSplitted := strings.Split(strNumbers, ",")
// // strNumbersSplitted := strings.Fields(strNumbers)
// // fmt.Println(strNumbersSplitted) 

// var strOnlyNumbers strings.Builder
// for _, value := range strNumbers {
// 	if( unicode.IsSpace(value) ){
// 		continue
// 	} else {
// 		strOnlyNumbers.WriteRune(value)
// 	}
// }

// result := strOnlyNumbers.String()

// fmt.Println(result) 
// // splNumbers := strings.Split(strOnlyNumbers, ",")

// }



func readingUsersField() (string, []float64, error){

    var operation string
    fmt.Print("Введите операцию (avg/sum/med): ")
    fmt.Scanln(&operation)
	operation = strings.ToLower(operation)


    fmt.Print("Введите числа через запятую: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()


    cleaned := strings.ReplaceAll(input, " ", "")
	strNumbers := strings.Split(cleaned, ",")
    

    var numbers []float64
    for _, s := range strNumbers {
        if s == "" {
            continue
        }
        num, err := strconv.ParseFloat(s, 64)
        if err != nil {
            fmt.Printf("Ошибка: '%s' не является числом\n", s)
            return "", nil, err
        }
        numbers = append(numbers, num)
    }

return operation, numbers, nil
}

func calculation (operation string, numbers []float64) float64{
var result float64
sort.Float64s(numbers)

	switch operation {
	case "avg":
    for _, value := range numbers {
        result += value
    }
    result /= float64(len(numbers))

	case "sum":
		for _, value := range numbers {
			result += value
		}

	case "med":
		length := len(numbers) % 2
		if length == 0 {
			result = float64( (len(numbers) + 1)/ 2)
		} else {
			result = float64( len(numbers) / 2 +  (len(numbers) / 2 + 1) / 2)
		}

	}

	return result
}