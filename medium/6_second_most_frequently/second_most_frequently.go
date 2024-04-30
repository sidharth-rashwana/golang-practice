package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func secondHighestChar(s string) {
	m := make(map[rune]int)

	for _, ch := range s {
		m[ch]++
	}

	highest, second := 0, 0

	for _, count := range m {
		if count > highest {
			second = highest
			highest = count
		} else if count > second && count < highest {
			second = count
		}
	}

	fmt.Println("Frequencies:", m)
	fmt.Print("Second highest characters: ")

	for ch, count := range m {
		if count == second {
			fmt.Printf("%c ", ch)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Enter a string: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	secondHighestChar(s)
}
