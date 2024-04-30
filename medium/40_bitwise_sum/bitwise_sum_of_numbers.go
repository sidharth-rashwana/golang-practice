package main

import "fmt"

func getSum(a, b int) int {
	for b != 0 {
		// Calculate the sum without carry
		tempSum := a ^ b
		// Calculate the carry and shift it left
		carry := (a & b) << 1
		// Update a and b for the next iteration
		a = tempSum
		b = carry
	}
	return a
}

func main() {
	fmt.Println(getSum(2, 3))
	fmt.Println(getSum(8, 7))
}
