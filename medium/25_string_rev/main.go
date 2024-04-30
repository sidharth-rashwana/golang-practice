package main

import "fmt"

func main() {

	s := "abcde"
	rs := []rune(s)
	left, right := 0, len(rs)-1
	for left < right {
		temp := rs[left]
		rs[left] = rs[right]
		rs[right] = temp
		left++
		right--
	}
	fmt.Println(string(rs))
}
