package main

import (
	"fmt"
	"sort"
)

func findSmallestMissingPositive(orderNumbers []int32) int32 {
	// Sort the slice safely
	sort.Slice(orderNumbers, func(i, j int) bool {
		return orderNumbers[i] < orderNumbers[j]
	})

	fmt.Println("ascending order: ", orderNumbers)

	smallest := int32(1)

	for _, num := range orderNumbers {
		if num == smallest {
			// only increase if it matches
			smallest++
		}
	}

	return smallest
}

func main() {
	fmt.Println("smallest missing positive number : ", findSmallestMissingPositive([]int32{3, 4, -1, 1}))
}
