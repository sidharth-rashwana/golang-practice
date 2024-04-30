package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isAlphabeticPalindrome(code string) bool {
	lowerCode := strings.ToLower(code)
	ch := []string{}
	for idx := 0; idx < len(lowerCode); idx++ {
		isLetter := unicode.IsLetter(rune(lowerCode[idx]))
		if isLetter {
			ch = append(ch, string(lowerCode[idx]))
		}
	}
	cur := len(ch) - 1
	for idx := 0; idx < len(ch); idx++ {
		if string(ch[idx]) != string(ch[cur]) {
			return false
		}
		if string(ch[idx]) == string(ch[cur]) {
			cur--
		}
	}
	return true
}

func main() {
	fmt.Println(isAlphabeticPalindrome("A1b2B!a"))
}
