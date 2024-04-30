package main

import "fmt"

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Push(x int) {
	h.arr = append(h.arr, x)
	i := len(h.arr) - 1

	for i > 0 {
		p := (i - 1) / 2
		if h.arr[p] <= h.arr[i] {
			break
		}
		h.arr[p], h.arr[i] = h.arr[i], h.arr[p]
		i = p
	}
}

func (h *MinHeap) Pop() int {
	if len(h.arr) == 0 {
		return -1
	}
	min := h.arr[0]
	last := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	if len(h.arr) > 0 {
		h.arr[0] = last
		h.heapify(0)
	}
	return min
}

func (h *MinHeap) heapify(i int) {
	smallest := i
	l, r := 2*i+1, 2*i+2

	if l < len(h.arr) && h.arr[l] < h.arr[smallest] {
		smallest = l
	}
	if r < len(h.arr) && h.arr[r] < h.arr[smallest] {
		smallest = r
	}
	if smallest != i {
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		h.heapify(smallest)
	}
}

func main() {
	h := MinHeap{}
	h.Push(3)
	h.Push(1)
	h.Push(2)

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
}
