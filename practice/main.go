package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Time now: ", n)

	t := time.Date(2022, time.April, 8, 10, 37, 0, 0, time.UTC)
	fmt.Println("Time date: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Apr  8 10:37:00 2022")
	fmt.Printf("Parsed time type: %T\n", parsedTime)
}
