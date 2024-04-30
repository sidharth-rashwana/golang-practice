package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longestShortest(s string) (int, int) {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0, 0
	}

	longest := len(words[0])
	shortest := len(words[0])

	for _, w := range words {
		if len(w) > longest {
			longest = len(w)
		}
		if len(w) < shortest {
			shortest = len(w)
		}
	}

	return longest, shortest
}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(longestShortest(s))
}
