package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	input1 := readInput("Value 1: ")
	float1 := convToFloat64(input1)

	input2 := readInput("Value 2: ")
	float2 := convToFloat64(input2)

	operator := readInput("Select an operation (+ - / *): ")

	calcVal := performCalc(float1, float2, operator)

	fmt.Printf("The result is %v\n\n", calcVal)
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func convToFloat64(input string) float64 {
	floatVal, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	return floatVal
}

func performCalc(num1, num2 float64, operator string) float64 {
	// matchOperator(operator)

	var retVal float64

	switch operator {
	case "+":
		retVal = num1 + num2
	case "-":
		retVal = num1 - num2
	case "/":
		retVal = num1 / num2
	case "*":
		retVal = num1 * num2
	default:
		panic("Invalid operation")
	}

	return math.Round(retVal*100) / 100
}

// func matchOperator(operator string) {
// 	allowedOperators := []string{"+", "-", "/", "*"}

// 	if !contains(allowedOperators, operator) {
// 		panic("Operator does not match + - / *")
// 	}
// }

// func contains(s []string, str string) bool {
// 	for _, v := range s {
// 		if str == v {
// 			return true
// 		}
// 	}
// 	return false
// }
