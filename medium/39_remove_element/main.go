/*Given set of numbers d = [1,2,3,1,5] , remove all occurance of target = [1] */

package main

import "fmt"

func removeElement(nums []int, val int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[c] = nums[i]
			c++
		}
	}
	return c
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
