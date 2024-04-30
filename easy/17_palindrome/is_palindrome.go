// a sentence or phrase that can be read same backward and forward
package main

import "fmt"

func isPalindrome(s string) bool {
	for idx := 0; idx < len(s); idx++ {
		if s[idx] != s[len(s)-idx-1] { // base condition
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Println("Enter a string : ")
	fmt.Scanln(&input)
	fmt.Println(isPalindrome(input))
}
