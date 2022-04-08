package main

import (
	"fmt"
	"math/rand"
	"sort"
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
		numbers[i] = rand.Intn(100)
	}

	numbers = append(numbers, 40)

	fmt.Println(numbers)

	sort.Ints(numbers)

	fmt.Println(numbers)
}
