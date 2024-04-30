package main

import (
	"fmt"
	"unicode"
)

/*
Check for Non-Identical String Rotation
*/
func isNonTrivialRotation(s1 string, s2 string) bool {
	// base case
	if len(s1) != len(s2) || s1 == s2{
		return false
	}
	for _, val := range s1 {
		if !unicode.IsLower(val) {
			return false
		}
	}

	for _, val := range s2 {
		if !unicode.IsLower(val) {
			return false
		}
	}

	for idx := 1; idx < len(s1); idx++ {
		cmp := s1[idx:] + s1[:idx]
		fmt.Println("cmp", cmp)
		if cmp == s2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isNonTrivialRotation("abcde", "cdeab"))
}
