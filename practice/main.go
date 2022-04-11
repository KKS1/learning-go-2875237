package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	fmt.Println(content)

	tours := toursFromJson(content)

	for _, tour := range tours {
		price, err := strconv.ParseFloat(strings.TrimSpace(tour.Price), 64)
		checkError(err)
		fmt.Printf("Name: %v, Price: %v \n", tour.Name, price)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	checkError(err)

	var tour Tour

	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
