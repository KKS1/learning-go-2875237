package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for index, v := range colors {
		fmt.Println(index, v)
	}

	val := 1

	for val < 5 {
		fmt.Println(val)
		val++
		if val > 3 {
			goto FOR_LOOP_END
		}
	}

FOR_LOOP_END:
	fmt.Println("For loop ending statement")
}
