//Write a Go function to count the number of words in a string and display count for each word.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordMap(words []string) map[string]int {
	m := make(map[string]int)
	for _, val := range words {
		m[val]++
	}
	return m

}

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	fmt.Println(wordMap(words))
}
