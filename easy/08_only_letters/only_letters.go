// Write a Go function to check if a string contains only letters.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func onlyLettersM1(s string) bool {
	a := []rune{}
	for i := 'a'; i <= 'z'; i++ {
		a = append(a, i)
	}
	
	for j := 'A'; j <= 'Z'; j++ {
		a = append(a, j)
	}

	for _, value := range s {
		if !strings.Contains(string(a), string(value)) {
			return false
		}

	}
	return true
}

func onlyLettersM2(s string) bool {
	for _, value := range s {
		if !unicode.IsLetter(value) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Enter a string: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(onlyLettersM1(s))
	fmt.Println(onlyLettersM2(s))
}
