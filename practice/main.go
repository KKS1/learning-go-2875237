package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["CA"] = "California"
	states["NJ"] = "New Jersey"
	states["NY"] = "New York"
	fmt.Println(states)
	fmt.Println(states["CA"])

	delete(states, "NY")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)

	sort.Strings(keys)

	fmt.Println(keys)

	for j := range keys {
		fmt.Println(states[keys[j]])
	}
}
