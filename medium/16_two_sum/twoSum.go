package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Map to store value -> index
	numMap := make(map[int]int)

	for i, val := range nums {
		// Check if the complement exists in the map
		if j, ok := numMap[target-val]; ok {
			return []int{j, i}
		}
		// Store the current value and its index
		numMap[val] = i // we are storing index as value
	}

	return []int{}
}

func main() {
	x := []int{2, 7, 11, 15}
	fmt.Println(twoSum(x, 9)) // [0 1]

	y := []int{3, 2, 4}
	fmt.Println(twoSum(y, 6)) // [1 2]

	z := []int{3, 3}
	fmt.Println(twoSum(z, 6)) // [0 1]
}
