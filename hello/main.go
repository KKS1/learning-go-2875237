package main

import (
	"fmt"
)

/*
In the following example, we will compute the sum of numbers in an array.
We split the array among two goroutines and find the sum of the two array slices concurrently.
After finding sum of the slices, we write the values to channel.
In the main goroutine, we will receive the value from the channel and find the complete sum.
*/

var myChannel = make(chan int)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	mid := (len(arr) / 2)

	splitArr1 := arr[:mid]

	splitArr2 := arr[mid:]

	go sum(splitArr1)

	go sum(splitArr2)

	// chanLen := len(myChannel)
	// fmt.Println("Channel length: ", chanLen)

	sum1 := <-myChannel
	sum2 := <-myChannel

	sum := sum1 + sum2

	fmt.Printf("Sum of %v and %v is: %v\n", sum1, sum2, sum)
}

func sum(vals []int) {
	_sum := 0
	for _, v := range vals {
		_sum += v
	}
	myChannel <- _sum
}
