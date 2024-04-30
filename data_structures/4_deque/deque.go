package deque

import "fmt"

type Deque struct {
	data []int
}

func (d *Deque) PushFront(x int) {
	d.data = append([]int{x}, d.data...)
}

func (d *Deque) PushBack(x int) {
	d.data = append(d.data, x)
}

func (d *Deque) PopFront() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

func (d *Deque) PopBack() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return val
}

func main() {
	d := Deque{}
	d.PushFront(1)
	d.PushBack(2)
	d.PushFront(0)

	fmt.Println(d.PopFront())
	fmt.Println(d.PopBack())
	fmt.Println(d.PopFront())
}
