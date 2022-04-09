package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic("Please enter a number")
	}

	fmt.Printf("You entered: %v\n", num)

	result := ""

	if num < 0 {
		result = "less than zero"
	} else if num == 0 {
		result = "equals zero"
	} else {
		result = "greater than zero"
	}

	fmt.Println(result)
}
