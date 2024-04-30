//Write a Go function to check if a string contains only digits.

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

func onlyDigits(s string) bool {
	for _, val := range s {
		if !unicode.IsDigit(val) {
			return false
		}
	}
	return true
}

// return once find 1st occurance of a digit
func is_digit(x string) {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, val := range x {
		if slices.Contains(digits, val) {
			fmt.Printf("This string '%s' does contain a digit : '%v' ", x, string(val))
			return
		}
	}
	fmt.Printf("This string %s does not contain a digit. ", x)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(onlyDigits(s))

	is_digit("this is a sample string : 12345")
	is_digit("this is a sample string")
}
