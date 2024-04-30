package main

import (
	"fmt"
)

func searchRotatedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // Target found
		}

		// Check which half is sorted
		if nums[left] <= nums[mid] { // Left half is sorted
			if nums[left] <= target && target < nums[mid] { // Target in left half
				right = mid - 1
			} else { // Target in right half
				left = mid + 1
			}
		} else { // Right half is sorted
			if nums[mid] < target && target <= nums[right] { // Target in right half
				left = mid + 1
			} else { // Target in left half
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	arr := []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	target := 3
	index := searchRotatedArray(arr, target)
	fmt.Println("Index of target:", index)
}
