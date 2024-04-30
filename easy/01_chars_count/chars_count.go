// Write a Go function to count the number of characters in a string.
package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	m := make(map[string]int)

	for _, val := range s {
		m[string(val)]++
	}

	fmt.Println(m)
}
