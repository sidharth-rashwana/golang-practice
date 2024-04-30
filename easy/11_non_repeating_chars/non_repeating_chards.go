package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	d := make(map[string]int)
	for _, val := range str {
		d[string(val)]++
		if d[string(val)] > 1 {
			delete(d, string(val))
		}
	}

	// for k, val := range d {
	// 	if val > 1 {
	// 		delete(d, k)
	// 	}
	// }

	fmt.Println(d)
}
