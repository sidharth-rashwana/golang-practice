package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sample map
	myMap := map[string]int{
		"banana": 2,
		"apple":  5,
		"cherry": 1,
	}

	// Step 1: Extract the keys
	var keys []string
	for k := range myMap {
		keys = append(keys, k)
	}

	// Step 2: Sort the keys
	sort.Strings(keys)

	// Step 3: Access the map in sorted order
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, myMap[k])
	}
}
