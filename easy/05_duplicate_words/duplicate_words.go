package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findDuplicates(s string) string {
	m := make(map[string]int)
	w := strings.Fields(s)
	for _, val := range w {
		m[string(val)]++
		if m[string(val)] > 1 {
			return string(val)
		}
	}
	return "None are duplicate"
}

func main() {
	var input string
	fmt.Println("Enter a string")
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(findDuplicates(input))
}
