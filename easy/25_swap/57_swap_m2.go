package main

import "fmt"

func main() {
	var (
		input1, input2 int
	)
	fmt.Scanf("%d", &input1)
	fmt.Scanf("%d", &input2)

	input1, input2 = input2, input1
	fmt.Println(input1, input2)

	y := &input1
	z := &input2

	fmt.Printf("%p\n", y)
	fmt.Printf("%p\n", z)
}
