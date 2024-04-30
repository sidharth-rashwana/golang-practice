/*Create a function to generate the Fibonacci sequence up to n terms.*/
package main

import "fmt"

func Fibonacci(n int) {

	first, second := 0, 1
	var next int
	for i := 0; i < n; i++ {
		print(first, "\n")
		next = first + second
		first = second
		second = next
	}

}

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)
	Fibonacci(n)
}
