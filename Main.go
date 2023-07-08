package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	readLine, _ := reader.ReadString('\n')
	stringWithoutSpace := strings.ReplaceAll(readLine, " ", "")

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*/,_+-\r\n", r)
	}
	words := strings.FieldsFunc(stringWithoutSpace, splitFunc)

	if len(words) > 2 {
		err := errors.New("вы ввели более двух чисел")
		fmt.Println("ОШИБКА:", err)
	} else if len(words) < 2 {
		err := errors.New("вы ввели менее двух чисел")
		fmt.Println("ОШИБКА:", err)
	}

	romeNumbers := [...]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	Numbers := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	num1 := words[0]
	if !(numberInArray(num1, romeNumbers) || numberInArray(num1, Numbers)) {
		err := errors.New("Первое число > 10")
		fmt.Println("ОШИБКА:", err)
	}

	num2 := words[1]

	if !(numberInArray(num2, romeNumbers) || numberInArray(num2, Numbers)) {
		err := errors.New("Второе число > 10")
		fmt.Println("ОШИБКА:", err)
	}

	splitFunc = func(r rune) bool {
		return strings.ContainsRune(num1+num2, r)
	}
	words2 := strings.FieldsFunc(stringWithoutSpace, splitFunc)

	symbol := words2[0]

	if numberInArray(num1, romeNumbers) && numberInArray(num2, romeNumbers) {

		romeNum1 := stringToNum(num1, romeNumbers)
		romeNum2 := stringToNum(num2, romeNumbers)

		resultOfOperation := operation(romeNum1, symbol, romeNum2)
		if resultOfOperation < 1 {
			err := errors.New("в Римской системе нет отрицательных чисел")
			fmt.Println("ОШИБКА:", err)
		} else {
			romeResult := simpleInRome(resultOfOperation, romeNumbers)

			fmt.Println(romeResult)
		}
	}
	if numberInArray(num1, Numbers) && numberInArray(num2, Numbers) {
		SimpleNum1 := stringToNum(num1, Numbers)
		SimpleNum2 := stringToNum(num2, Numbers)

		resultOfOperation := operation(SimpleNum1, symbol, SimpleNum2)
		fmt.Println(resultOfOperation)
	}
	if numberInArray(num1, romeNumbers) && numberInArray(num2, Numbers) || numberInArray(num2, romeNumbers) && numberInArray(num1, Numbers) {
		err := errors.New("Нельзя использовать разные системы исчисления")
		fmt.Println("ОШИБКА:", err)
	}

}
func operation(num1 int, symbol string, num2 int) int {
	var result int
	if symbol == "*" {
		result = num1 * num2
	} else if symbol == "/" {
		result = num1 / num2
	} else if symbol == "+" {
		result = num1 + num2
	} else if symbol == "-" {
		result = num1 - num2
	}
	return result
}
func numberInArray(num string, arr [10]string) bool {
	var result bool
	for _, value := range arr {
		if value == num {
			result = true
			break
		}
	}
	return result
}
func stringToNum(num string, arr [10]string) int {
	var romeNum int
	for key, value := range arr {
		if value == num {
			romeNum = key + 1
			break
		}
	}
	return romeNum
}
func simpleInRome(num int, arr [10]string) string {
	var number string = ""
	if num%100 == 0 {
		number = "C"
	} else if num%50 == 0 {
		number = "L"
	} else if num/50 > 0 {
		number0 := "L"
		result := num % 50
		if result > 10 {
			number1 := ""
			xresult := result / 10
			for i := 0; i < xresult; i++ {
				number1 = number1 + "X"
			}
			result2 := result % 10
			number2 := arr[result2-1]
			number = number0 + number1 + number2
		} else if result < 10 {
			number1 := arr[result-1]
			number = number0 + number1
		} else if result == 10 {
			number1 := "X"
			number = number0 + number1
		}
	} else {
		if num < 10 {
			number = arr[num-1]
		} else if num > 10 {
			number1 := ""
			xresult := num / 10
			for i := 0; i < xresult; i++ {
				number1 = number1 + "X"
			}
			result := num % 10
			if result > 0 {
				number2 := arr[result-1]
				number = number1 + number2
			} else {
				number = number1
			}

		}
	}
	return number
}
