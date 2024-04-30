package main

import "fmt"

/*
What it should compute:

You want to find the maximum area that can be formed between two vertical lines.

The area is determined by:

	area=min(height[i],height[j])×(j−i)

1. min(height[i], height[j]) → because the water is limited by the shorter wall.

2. (j - i) → because the width of the container is the distance between the two lines.

*/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxOutput := 0
	for left < right {
		h := min(height[left], height[right])
		width := right - left
		area := h * width

		if area > maxOutput {
			maxOutput = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxOutput
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
