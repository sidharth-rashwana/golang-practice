package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Implement a stack that supports push, pop, top, and getMin operations in O(1) time, where getMin returns the minimum element.

# Example

# Input

n = 10
operations = ['push 2', 'push 0', 'push 3', 'push 0', 'getMin', 'pop', 'getMin', 'pop', 'top', 'getMin']
Output

[0,0,0,0]
Explanation

- push 2 → stack = [2], min = 2 2.
- push 0 → stack = [2,0], min = 0 3.
- push 3 → stack = [2,0,3], min = 0 4.
- push 0 → stack = [2,0,3,0], min = 0 5.
- getMin → returns 0 6.
- pop → removes 0, stack = [2,0,3], min = 0 7.
- getMin → returns 0 8.
- pop → removes 3, stack = [2,0], min = 0 9.
- top → returns 0 10. getMin → returns 0

*/
func processCouponStackOperations(operations []string) []int32 {
	var stack []int32
	var minStack []int32
	var result []int32

	for _, op := range operations {

		if strings.HasPrefix(op, "push") {
			// extract number after "push "
			parts := strings.Split(op, " ")
			val, _ := strconv.Atoi(parts[1])
			v := int32(val)

			stack = append(stack, v)

			if len(minStack) == 0 || v <= minStack[len(minStack)-1] {
				minStack = append(minStack, v)
			} else {
				minStack = append(minStack, minStack[len(minStack)-1])
			}

		} else if op == "pop" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				minStack = minStack[:len(minStack)-1]
			}

		} else if op == "top" {
			if len(stack) > 0 {
				result = append(result, stack[len(stack)-1])
			}

		} else if op == "getMin" {
			if len(minStack) > 0 {
				result = append(result, minStack[len(minStack)-1])
			}
		}
	}

	return result
}

func main() {
	operations := []string{"push 2", "push 0", "push 3", "push 0", "getMin", "pop", "getMin", "pop", "top", "getMin"}
	fmt.Println(processCouponStackOperations(operations))

}
