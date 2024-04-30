package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 4, 5, 6, 7}

	result := make([]int, 0, len(x)+len(y))

	i, j := 0, 0

	for i < len(x) && j < len(y) {
		if x[i] <= y[j] {
			result = append(result, x[i])
			i++
		} else {
			result = append(result, y[j])
			j++
		}
	}
	// Append remaining elements
	result = append(result, x[i:]...)
	result = append(result, y[j:]...)

	fmt.Println(result)
}
