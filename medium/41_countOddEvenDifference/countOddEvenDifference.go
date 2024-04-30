package main

import "fmt"

func countOddEvenDifference(l []int) {

	even, odd := 0, 0

	for _, val := range l {

		if val%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println(even - odd)
}

func main() {

	countOddEvenDifference([]int{1, 2, 3, 4, 5, 6, 7})
}
