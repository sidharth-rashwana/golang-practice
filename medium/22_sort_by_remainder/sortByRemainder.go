/*
1. sort by modulo
2. sort alphabetically if modulo is same
*/
package main

import (
	"fmt"
	"sort"
)

func sortStable(s []string) []string {
	sort.SliceStable(s, func(i, j int) bool {
		fmt.Println(s[i], s[j])
		ri := len(s[i]) % 3
		rj := len(s[j]) % 3

		if ri != rj {
			return ri < rj
		}

		return s[i] < s[j]
	})
	return s

}

func main() {
	s := []string{"a", "ab", "bc", "abc"}
	fmt.Println(sortStable(s)) // [abc a ab bc]
}
