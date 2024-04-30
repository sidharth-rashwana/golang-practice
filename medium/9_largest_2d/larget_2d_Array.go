// Implement a program to find the largest element in a 2D array.
// here largest element is '9'

package main

import "fmt"

func largestElement(arr2d [][]int) int {
	largest := arr2d[0][0]

	for _, row := range arr2d {
		fmt.Println(row)
		for _, val := range row {
			print(val)
			if val > largest {
				largest = val
			}
		}
	}

	return largest
}

func main() {
	var arr2d = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	l := largestElement(arr2d)
	fmt.Println(l)
}
