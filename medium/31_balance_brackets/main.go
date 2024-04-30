package main

import "fmt"

func areBracketsProperlyMatched(code_snippet string) bool {
	if len(code_snippet) == 0 {
		return false
	}
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range code_snippet {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0

}

func main() {
	fmt.Println("is balanced: ", areBracketsProperlyMatched("if (a[0] > b[1]) { doSomething(); }"))
}
