package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("A simple calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')

	value1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		fmt.Println(err1)
		panic("Please enter a number")
	}

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')

	value2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err2 != nil {
		fmt.Println(err2)
		panic("Please enter a number")
	}

	fmt.Println("You entered: ", value2)

	finalSum := value1 + value2
	finalSum = math.Round(finalSum*100) / 100
	fmt.Printf("Final Sum: %v\n", finalSum)
}
