// https://leetcode.com/problems/container-with-most-water/
package main

import "fmt"

// maxArea function to find the container that holds the most water
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	// Loop until the two pointers meet
	for left < right {
		// Calculate the area
		width := right - left
		currentArea := min(height[left], height[right]) * width
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// Helper function to get the minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example test case
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	fmt.Println("Maximum water that can be contained:", result)
}
