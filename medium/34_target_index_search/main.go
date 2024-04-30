package main

import "fmt"

func binarySearch(nums []int32, target int32) int32 {
	for idx, val := range nums {
		if int32(val) == target {
			return int32(idx)
		}
	}
	return -1
}

func main() {
	fmt.Println("index: ", binarySearch([]int32{1, 2, 3, 4, 5}, 5))
}
