package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Printf("%v\n", poodle)
	poodle.Weight = 11
	fmt.Printf("%v\n", poodle)
}

type Dog struct {
	Breed  string
	Weight int
}
