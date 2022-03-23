package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum: ", intSum)

	f1, f2, f3 := 1.1233, 2.2211, 3.3133
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum: ", floatSum)

	roundedSum := math.Round(floatSum*100) / 100
	fmt.Println("Rounded Sum: ", roundedSum)

	radius := 2.12
	circumference := 2 * math.Pi * radius
	fmt.Printf("Circumference is %.2f\n", circumference)
}
