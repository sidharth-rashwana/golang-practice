package main

import (
	"fmt"
	"strings"
	"unicode"
)

func convUpperM1(s string) string {
	x := ""
	for _, val := range s {
		if !unicode.IsUpper(val) {
			x += strings.ToUpper(string(val))
		} else {
			x += string(val)
		}
	}
	return x
}

func convUpperM2(s string) string {
	return strings.ToUpper(s)
}

func main() {
	var x string
	fmt.Scanln(&x)
	fmt.Println(convUpperM1(x))
	fmt.Println(convUpperM2(x))
}
