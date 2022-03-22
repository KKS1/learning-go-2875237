package main

import (
	"fmt"
)

const abc = "custom constant"

func main() {
	var myStr string = "some string here"
	smStr := "Hello from Go!"
	fmt.Println(smStr, myStr, abc)
	fmt.Printf("variable types=> smStr: %T , myStr: %T , abc: %T\n", smStr, myStr, abc)
}
