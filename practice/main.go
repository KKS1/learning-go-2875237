package main

import (
	"fmt"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)

	for i := 0; i < 5; i++ {
		numbers[i] = 134 + i
	}

	numbers = append(numbers, 140)

	fmt.Println(numbers)
}
