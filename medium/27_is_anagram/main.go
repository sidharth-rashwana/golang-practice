package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) int32 {
	s = strings.ToLower(s)
	t = strings.ToLower(t)
	if len(s) != len(t) {
		return 0
	}

	s1 := []rune(s)
	s2 := []rune(t)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})

	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	if string(s1) == string(s2) {
		return 1
	}
	return 0
}

func main() {
	fmt.Println("is anagram: 1 --> yes, 0 --> no", isAnagram("listen", "tenlis"))
}
