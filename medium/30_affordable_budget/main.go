package main

import "fmt"

func countAffordablePairs(prices []int32, budget int32) int32 {

	left, right := 0, len(prices)-1
	pairs := int32(0)

	for left < right {
		if prices[left]+prices[right] <= budget {
			pairs += int32(right - left)
			left++
		} else {
			right--
		}
	}

	return int32(pairs)
}

func main() {
	fmt.Println("affordable budget : ", countAffordablePairs([]int32{1, 2, 3, 4, 5}, 7))
}
