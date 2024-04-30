package main

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1 // Minimum is in the right half
		} else {
			right = mid // Minimum is in the left half (including mid)
		}
	}

	return nums[left] // Left now points to the minimum
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // Output: 0
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // Output: 1
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // Output: 11
	fmt.Println(findMin([]int{2, 1, 0}))             // Output: 1
}
