package main

import (
	"fmt"
)

func main() {
	doSomething()
	twoSum := addTwoValues(1, 2)
	fmt.Println("Two values sum: ", twoSum)

	allSum, lenValues := addAllValues(1, 2, 3, 4, 5)
	fmt.Printf("All values sum and length: %v %v\n", allSum, lenValues)
}

func doSomething() {
	fmt.Println("Functions")
}

func addTwoValues(val1, val2 int) int {
	return val1 + val2
}

func addAllValues(values ...int) (int, int) {
	total := 0

	for _, val := range values {
		total = total + val
	}

	return total, len(values)
}
