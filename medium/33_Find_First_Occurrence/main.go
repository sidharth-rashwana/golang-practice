package main

import "fmt"

func findFirstOccurrence(nums []int32, target int32) int32 {
	low, high := 0, len(nums)-1
	result := int32(-1)

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			result = int32(mid)
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

func main() {
	fmt.Println(findFirstOccurrence([]int32{1, 2, 3, 3, 4, 5}, 3)) // Output: 2
}
