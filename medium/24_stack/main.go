package main

import (
	"fmt"
	"os"
)

var visited = []int32{}

type Stack struct {
}

func (s Stack) Push(x int32) {
	visited = append(visited, x)
}

func (s Stack) Pop() {
	// if stack length is already 0 , you cannot pop any element
	if len(visited) == 0 {
		fmt.Println("No elements in the stack ...")
		return
	}
	visited = visited[:len(visited)-1]
	return
}

func (s Stack) Display() {
	for idx := 0; idx < len(visited); idx++ {
		fmt.Println(visited[idx])
	}
}

func main() {
	var s Stack

	for {
		fmt.Println("Enter your choice : 1. Push 2. Pop 3. Display 4. Exit")
		var operation string
		fmt.Scanf("%s", &operation)

		switch operation {
		case "1":
			var num int32
			fmt.Println("enter a number : ")
			fmt.Scanf("%d", &num)
			s.Push(num)

		case "2":
			s.Pop()

		case "3":
			fmt.Println("Display stack : ")
			s.Display()

		default:
			fmt.Println("None case matched ... Exiting !!!")
			os.Exit(200)
		}

	}
}
