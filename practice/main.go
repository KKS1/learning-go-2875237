package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)
	checkError(err)

	fmt.Printf("Response type: %T \n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)

	fmt.Println(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
