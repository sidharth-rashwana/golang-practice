package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func recurring_char(s string) {
	m := make(map[string]int)
	for _, val := range s {
		m[string(val)]++
		if m[string(val)] > 1 {
			fmt.Println("recurring", m[string(val)])
		}
	}

}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	recurring_char(s)
}
