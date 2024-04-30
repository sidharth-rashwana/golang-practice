// Write a Go function to reverse a string.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	// swap
	runes := []rune(s) //convert string into runes
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func reverseStringM2(s string) string {
	// swap
	runes := []rune(s) //convert string into runes
	left, right := 0, len(runes)-1
	for left < right {
		temp := runes[left]
		runes[left] = runes[right]
		runes[right] = temp
		left++
		right--
	}
	return string(runes)
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(reverseString(s))
	fmt.Println(reverseStringM2(s))
}
